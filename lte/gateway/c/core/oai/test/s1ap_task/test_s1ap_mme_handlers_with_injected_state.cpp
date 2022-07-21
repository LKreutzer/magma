#include <gtest/gtest.h>
#include <thread>

#include "lte/gateway/c/core/oai/test/mock_tasks/mock_tasks.hpp"

extern "C" {
#include "lte/gateway/c/core/common/dynamic_memory_check.h"
#include "lte/gateway/c/core/oai/common/log.h"
#include "lte/gateway/c/core/oai/include/mme_config.h"
#include "lte/gateway/c/core/oai/lib/bstr/bstrlib.h"
#include "lte/gateway/c/core/oai/include/mme_init.hpp"
}

#include "lte/gateway/c/core/oai/include/s1ap_state.hpp"
#include "lte/gateway/c/core/oai/tasks/s1ap/s1ap_mme_decoder.hpp"
#include "lte/gateway/c/core/oai/tasks/s1ap/s1ap_mme_nas_procedures.hpp"
#include "lte/gateway/c/core/oai/tasks/s1ap/s1ap_mme_handlers.hpp"
#include "lte/gateway/c/core/oai/test/s1ap_task/s1ap_mme_test_utils.h"
#include "lte/gateway/c/core/oai/test/s1ap_task/mock_s1ap_op.h"
#include "lte/gateway/c/core/oai/tasks/s1ap/s1ap_state_manager.hpp"

extern bool hss_associated;
extern task_zmq_ctx_t task_zmq_ctx_mme;

namespace magma {
namespace lte {

extern task_zmq_ctx_t task_zmq_ctx_main_s1ap_with_injected_states;

static int handle_message(zloop_t* loop, zsock_t* reader, void* arg) {
  MessageDef* received_message_p = receive_msg(reader);

  switch (ITTI_MSG_ID(received_message_p)) {
    default: {
    } break;
  }

  itti_free_msg_content(received_message_p);
  free(received_message_p);
  return 0;
}

class S1apMmeHandlersWithInjectedStatesTest : public ::testing::Test {
  virtual void SetUp() {
    mme_app_handler = std::make_shared<MockMmeAppHandler>();
    sctp_handler = std::make_shared<MockSctpHandler>();

    itti_init(TASK_MAX, THREAD_MAX, MESSAGES_ID_MAX, tasks_info, messages_info,
              NULL, NULL);

    // initialize mme config
    mme_config_init(&mme_config);
    create_partial_lists(&mme_config);
    mme_config.use_stateless = false;
    hss_associated = true;

    task_id_t task_id_list[4] = {TASK_MME_APP, TASK_S1AP, TASK_SCTP,
                                 TASK_SERVICE303};
    init_task_context(TASK_MAIN, task_id_list, 4, handle_message,
                      &task_zmq_ctx_main_s1ap_with_injected_states);

    std::thread task_mme_app(start_mock_mme_app_task, mme_app_handler);
    std::thread task_sctp(start_mock_sctp_task, sctp_handler);
    task_mme_app.detach();
    task_sctp.detach();

    s1ap_mme_init(&mme_config);

    std::string magma_root = std::getenv("MAGMA_ROOT");
    std::string state_data_path = magma_root + "/" + DEFAULT_S1AP_STATE_DATA_PATH;
    std::string data_folder_path = magma_root + "/" + DEFAULT_S1AP_CONTEXT_DATA_PATH;
    std::string data_list_path = magma_root + "/" + DEFAULT_S1AP_CONTEXT_DATA_PATH + "data_list.txt";

    number_attached_ue = 2;

    mock_read_s1ap_state_db(state_data_path);
    name_of_ue_samples = load_file_into_vector_of_line_content(data_folder_path, data_list_path);
    mock_read_s1ap_ue_state_db(name_of_ue_samples);

    state = S1apStateManager::getInstance().get_state(false);
    assoc_id = 37;
    stream_id = 1;
  }

  virtual void TearDown() {
    // Sleep to ensure that messages are received and contexts are released
    std::this_thread::sleep_for(std::chrono::milliseconds(500));

    send_terminate_message_fatal(&task_zmq_ctx_main_s1ap_with_injected_states);
    send_terminate_message_fatal(&task_zmq_ctx_mme);

    destroy_task_context(&task_zmq_ctx_main_s1ap_with_injected_states);
    itti_free_desc_threads();

    free_mme_config(&mme_config);

    // Sleep to ensure that messages are received and contexts are released
    std::this_thread::sleep_for(std::chrono::milliseconds(200));
  }

 protected:
  std::shared_ptr<MockMmeAppHandler> mme_app_handler;
  std::shared_ptr<MockSctpHandler> sctp_handler;
  s1ap_state_t* state;
  sctp_assoc_id_t assoc_id;
  sctp_stream_id_t stream_id;
  std::vector<std::string> name_of_ue_samples;
  int number_attached_ue;
};

TEST_F(S1apMmeHandlersWithInjectedStatesTest, GenerateUEContextReleaseCommand) {
  ue_description_t ue_ref_p = {
      .enb_ue_s1ap_id = 1,
      .mme_ue_s1ap_id = 1,
      .sctp_assoc_id = assoc_id,
      .comp_s1ap_id = S1AP_GENERATE_COMP_S1AP_ID(assoc_id, 1)};
  ue_ref_p.s1ap_ue_context_rel_timer.id = -1;
  ue_ref_p.s1ap_ue_context_rel_timer.msec = 1000;

  ASSERT_EQ(RETURNok, s1ap_mme_generate_ue_context_release_command(
                          state, &ue_ref_p, S1AP_INITIAL_CONTEXT_SETUP_FAILED,
                          INVALID_IMSI64, assoc_id, stream_id, 1, 1));
}

}  // namespace lte
}  // namespace magma