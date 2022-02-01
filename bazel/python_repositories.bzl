# Copyright 2022 The Magma Authors.

# This source code is licensed under the BSD-style license found in the
# LICENSE file in the root directory of this source tree.

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""Hermetic Python interpreter configuration"""

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

PY_VERSION = "3.8.5"
PY_HASH = "e3003ed57db17e617acb382b0cade29a248c6026b1bd8aad1f976e9af66a83b0"
BUILD_DIR = "/tmp/bazel/external/python_{0}".format(PY_VERSION)

def _patch_cmds():
    return [
        "mkdir -p {0}".format(BUILD_DIR),
        "cp -r * {0}".format(BUILD_DIR),
        "cd {0} && ./configure --prefix={0}/bazel_install".format(BUILD_DIR),
        "cd {0} && make install".format(BUILD_DIR),
        "rm -rf * && mv {0}/* .".format(BUILD_DIR),
        "ln -s bazel_install/bin/python3 python_bin",
    ]
    
def python_repositories():

    http_archive(
        name = "python_interpreter",
        urls = ["https://www.python.org/ftp/python/{0}/Python-{0}.tar.xz".format(PY_VERSION)],
        sha256 = PY_HASH,
        strip_prefix = "Python-{0}".format(PY_VERSION),
        patch_cmds = _patch_cmds(),
        build_file = "//bazel/external:python_interpreter.BUILD",
    )
