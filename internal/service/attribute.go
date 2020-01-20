package service

import (
	"github.com/renatoaguimaraes/ecommerce-pim/internal/model"
	"github.com/renatoaguimaraes/ecommerce-pim/internal/repo"
)

// AttributeService service
type AttributeService struct {
	RepoAttr      repo.AttributeRepo
	RepoAttrValue repo.AttributeValueRepo
}

// InsertAttr a variant
func (vs AttributeService) InsertAttr(attr *model.Attribute) error {
	return vs.RepoAttr.Insert(attr)
}

// FindAttrByID a variant
func (vs AttributeService) FindAttrByID(group string, attributeID string) *model.Attribute {
	return vs.RepoAttr.FindAttributeByID(group, attributeID)
}

// InsertAttrValue a variant
func (vs AttributeService) InsertAttrValue(attrv *model.AttributeValue) error {
	return vs.RepoAttrValue.Insert(attrv)
}

// FindAttrValueByID updata a variant
func (vs AttributeService) FindAttrValueByID(attributeID string, attributeValueID string) *model.AttributeValue {
	return vs.RepoAttrValue.FindAttributeValueByID(attributeID, attributeValueID)
}
