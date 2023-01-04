// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	models3 "magma/orc8r/cloud/go/models"
	models4 "magma/orc8r/cloud/go/services/orchestrator/obsidian/models"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LteGateway Full description of an LTE gateway test change
//
// swagger:model lte_gateway
type LteGateway struct {

	// apn resources
	ApnResources ApnResources `json:"apn_resources,omitempty"`

	// cellular
	// Required: true
	Cellular *GatewayCellularConfigs `json:"cellular"`

	// checked in recently
	// Required: true
	CheckedInRecently *GatewayCheckedInRecently `json:"checked_in_recently"`

	// connected enodeb serials
	// Required: true
	ConnectedEnodebSerials EnodebSerials `json:"connected_enodeb_serials"`

	// description
	// Required: true
	Description models3.GatewayDescription `json:"description"`

	// device
	Device *models4.GatewayDevice `json:"device,omitempty"`

	// id
	// Required: true
	ID models3.GatewayID `json:"id"`

	// magmad
	// Required: true
	Magmad *models4.MagmadGatewayConfigs `json:"magmad"`

	// name
	// Required: true
	Name models3.GatewayName `json:"name"`

	// registration info
	RegistrationInfo *models3.RegistrationInfo `json:"registration_info,omitempty"`

	// status
	Status *models4.GatewayStatus `json:"status,omitempty"`

	// tier
	// Required: true
	Tier models4.TierID `json:"tier"`
}

// Validate validates this lte gateway
func (m *LteGateway) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApnResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCellular(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCheckedInRecently(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectedEnodebSerials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMagmad(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistrationInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTier(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LteGateway) validateApnResources(formats strfmt.Registry) error {
	if swag.IsZero(m.ApnResources) { // not required
		return nil
	}

	if m.ApnResources != nil {
		if err := m.ApnResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apn_resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apn_resources")
			}
			return err
		}
	}

	return nil
}

func (m *LteGateway) validateCellular(formats strfmt.Registry) error {

	if err := validate.Required("cellular", "body", m.Cellular); err != nil {
		return err
	}

	if m.Cellular != nil {
		if err := m.Cellular.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cellular")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cellular")
			}
			return err
		}
	}

	return nil
}

func (m *LteGateway) validateCheckedInRecently(formats strfmt.Registry) error {

	if err := validate.Required("checked_in_recently", "body", m.CheckedInRecently); err != nil {
		return err
	}

	if err := validate.Required("checked_in_recently", "body", m.CheckedInRecently); err != nil {
		return err
	}

	if m.CheckedInRecently != nil {
		if err := m.CheckedInRecently.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("checked_in_recently")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("checked_in_recently")
			}
			return err
		}
	}

	return nil
}

func (m *LteGateway) validateConnectedEnodebSerials(formats strfmt.Registry) error {

	if err := validate.Required("connected_enodeb_serials", "body", m.ConnectedEnodebSerials); err != nil {
		return err
	}

	if err := m.ConnectedEnodebSerials.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("connected_enodeb_serials")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("connected_enodeb_serials")
		}
		return err
	}

	return nil
}

func (m *LteGateway) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", models3.GatewayDescription(m.Description)); err != nil {
		return err
	}

	if err := m.Description.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("description")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("description")
		}
		return err
	}

	return nil
}

func (m *LteGateway) validateDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.Device) { // not required
		return nil
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *LteGateway) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", models3.GatewayID(m.ID)); err != nil {
		return err
	}

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *LteGateway) validateMagmad(formats strfmt.Registry) error {

	if err := validate.Required("magmad", "body", m.Magmad); err != nil {
		return err
	}

	if m.Magmad != nil {
		if err := m.Magmad.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("magmad")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("magmad")
			}
			return err
		}
	}

	return nil
}

func (m *LteGateway) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", models3.GatewayName(m.Name)); err != nil {
		return err
	}

	if err := m.Name.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("name")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("name")
		}
		return err
	}

	return nil
}

func (m *LteGateway) validateRegistrationInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.RegistrationInfo) { // not required
		return nil
	}

	if m.RegistrationInfo != nil {
		if err := m.RegistrationInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registration_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registration_info")
			}
			return err
		}
	}

	return nil
}

func (m *LteGateway) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *LteGateway) validateTier(formats strfmt.Registry) error {

	if err := validate.Required("tier", "body", models4.TierID(m.Tier)); err != nil {
		return err
	}

	if err := m.Tier.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tier")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("tier")
		}
		return err
	}

	return nil
}

// ContextValidate validate this lte gateway based on the context it is used
func (m *LteGateway) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApnResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCellular(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCheckedInRecently(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedEnodebSerials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMagmad(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegistrationInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTier(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LteGateway) contextValidateApnResources(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ApnResources.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("apn_resources")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("apn_resources")
		}
		return err
	}

	return nil
}

func (m *LteGateway) contextValidateCellular(ctx context.Context, formats strfmt.Registry) error {

	if m.Cellular != nil {
		if err := m.Cellular.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cellular")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cellular")
			}
			return err
		}
	}

	return nil
}

func (m *LteGateway) contextValidateCheckedInRecently(ctx context.Context, formats strfmt.Registry) error {

	if m.CheckedInRecently != nil {
		if err := m.CheckedInRecently.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("checked_in_recently")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("checked_in_recently")
			}
			return err
		}
	}

	return nil
}

func (m *LteGateway) contextValidateConnectedEnodebSerials(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ConnectedEnodebSerials.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("connected_enodeb_serials")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("connected_enodeb_serials")
		}
		return err
	}

	return nil
}

func (m *LteGateway) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Description.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("description")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("description")
		}
		return err
	}

	return nil
}

func (m *LteGateway) contextValidateDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.Device != nil {
		if err := m.Device.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *LteGateway) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *LteGateway) contextValidateMagmad(ctx context.Context, formats strfmt.Registry) error {

	if m.Magmad != nil {
		if err := m.Magmad.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("magmad")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("magmad")
			}
			return err
		}
	}

	return nil
}

func (m *LteGateway) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Name.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("name")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("name")
		}
		return err
	}

	return nil
}

func (m *LteGateway) contextValidateRegistrationInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.RegistrationInfo != nil {
		if err := m.RegistrationInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registration_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registration_info")
			}
			return err
		}
	}

	return nil
}

func (m *LteGateway) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *LteGateway) contextValidateTier(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Tier.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tier")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("tier")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LteGateway) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LteGateway) UnmarshalBinary(b []byte) error {
	var res LteGateway
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
