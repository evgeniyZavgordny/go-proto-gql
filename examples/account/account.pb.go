// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: account.proto

/*
Package account is a generated protocol buffer package.

It is generated from these files:
	account.proto

It has these top-level messages:
	SignInReq
	GetCurrentAccountReq
	SignUpWithEmailReq
	ResendConfirmationEmailReq
	SignOutReq
	ConfirmEmailReq
	ForgotPasswordReq
	CheckResetPasswordTokenReq
	RequestChangeEmailReq
	ChangeEmailReq
	ResetPasswordReq
	ChangePasswordReq
	RequestDeleteAccountReq
	DeleteAccountReq
	SignInRes
	SignUpWithEmailRes
	ConfirmEmailRes
	GetCurrentAccountRes
	CheckResetPasswordTokenRes
	SignOutRes
	ResendConfirmationEmailRes
	ForgotPasswordRes
	ResetPasswordRes
	ChangePasswordRes
	RequestChangeEmailRes
	ChangeEmailRes
	RequestDeleteAccountRes
	DeleteAccountRes
	Account
	SignedUpWithEmail
	EmailConfirmed
	PasswordReset
	PasswordChanged
	EmailChanged
	AccountDeleted
*/
package account

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/danielvladco/go-proto-gql"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Roles int32

const (
	Roles_DDD Roles = 0
	Roles_NNN Roles = 1
	Roles_GGG Roles = 2
)

var Roles_name = map[int32]string{
	0: "DDD",
	1: "NNN",
	2: "GGG",
}
var Roles_value = map[string]int32{
	"DDD": 0,
	"NNN": 1,
	"GGG": 2,
}

func (x Roles) String() string {
	return proto.EnumName(Roles_name, int32(x))
}
func (Roles) EnumDescriptor() ([]byte, []int) { return fileDescriptorAccount, []int{0} }

type ResendConfirmationEmailReq_Data int32

const (
	ResendConfirmationEmailReq_DDD ResendConfirmationEmailReq_Data = 0
	ResendConfirmationEmailReq_NNN ResendConfirmationEmailReq_Data = 1
	ResendConfirmationEmailReq_GGG ResendConfirmationEmailReq_Data = 2
)

var ResendConfirmationEmailReq_Data_name = map[int32]string{
	0: "DDD",
	1: "NNN",
	2: "GGG",
}
var ResendConfirmationEmailReq_Data_value = map[string]int32{
	"DDD": 0,
	"NNN": 1,
	"GGG": 2,
}

func (x ResendConfirmationEmailReq_Data) String() string {
	return proto.EnumName(ResendConfirmationEmailReq_Data_name, int32(x))
}
func (ResendConfirmationEmailReq_Data) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorAccount, []int{3, 0}
}

type ResendConfirmationEmailReq_Inner_InnerEnum int32

const (
	ResendConfirmationEmailReq_Inner_SSS ResendConfirmationEmailReq_Inner_InnerEnum = 0
	ResendConfirmationEmailReq_Inner_DDD ResendConfirmationEmailReq_Inner_InnerEnum = 1
)

var ResendConfirmationEmailReq_Inner_InnerEnum_name = map[int32]string{
	0: "SSS",
	1: "DDD",
}
var ResendConfirmationEmailReq_Inner_InnerEnum_value = map[string]int32{
	"SSS": 0,
	"DDD": 1,
}

func (x ResendConfirmationEmailReq_Inner_InnerEnum) String() string {
	return proto.EnumName(ResendConfirmationEmailReq_Inner_InnerEnum_name, int32(x))
}
func (ResendConfirmationEmailReq_Inner_InnerEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorAccount, []int{3, 0, 0}
}

type SignInReq struct {
	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (m *SignInReq) Reset()                    { *m = SignInReq{} }
func (m *SignInReq) String() string            { return proto.CompactTextString(m) }
func (*SignInReq) ProtoMessage()               {}
func (*SignInReq) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{0} }

func (m *SignInReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignInReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type GetCurrentAccountReq struct {
}

func (m *GetCurrentAccountReq) Reset()                    { *m = GetCurrentAccountReq{} }
func (m *GetCurrentAccountReq) String() string            { return proto.CompactTextString(m) }
func (*GetCurrentAccountReq) ProtoMessage()               {}
func (*GetCurrentAccountReq) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{1} }

type SignUpWithEmailReq struct {
	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (m *SignUpWithEmailReq) Reset()                    { *m = SignUpWithEmailReq{} }
func (m *SignUpWithEmailReq) String() string            { return proto.CompactTextString(m) }
func (*SignUpWithEmailReq) ProtoMessage()               {}
func (*SignUpWithEmailReq) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{2} }

func (m *SignUpWithEmailReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignUpWithEmailReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type ResendConfirmationEmailReq struct {
	Inner     ResendConfirmationEmailReq_Inner_InnerEnum `protobuf:"varint,4,opt,name=inner,proto3,enum=account.ResendConfirmationEmailReq_Inner_InnerEnum" json:"inner,omitempty"`
	Data      ResendConfirmationEmailReq_Data            `protobuf:"varint,2,opt,name=data,proto3,enum=account.ResendConfirmationEmailReq_Data" json:"data,omitempty"`
	AccountId string                                     `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (m *ResendConfirmationEmailReq) Reset()         { *m = ResendConfirmationEmailReq{} }
func (m *ResendConfirmationEmailReq) String() string { return proto.CompactTextString(m) }
func (*ResendConfirmationEmailReq) ProtoMessage()    {}
func (*ResendConfirmationEmailReq) Descriptor() ([]byte, []int) {
	return fileDescriptorAccount, []int{3}
}

func (m *ResendConfirmationEmailReq) GetInner() ResendConfirmationEmailReq_Inner_InnerEnum {
	if m != nil {
		return m.Inner
	}
	return ResendConfirmationEmailReq_Inner_SSS
}

func (m *ResendConfirmationEmailReq) GetData() ResendConfirmationEmailReq_Data {
	if m != nil {
		return m.Data
	}
	return ResendConfirmationEmailReq_DDD
}

func (m *ResendConfirmationEmailReq) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

type ResendConfirmationEmailReq_Inner struct {
}

func (m *ResendConfirmationEmailReq_Inner) Reset()         { *m = ResendConfirmationEmailReq_Inner{} }
func (m *ResendConfirmationEmailReq_Inner) String() string { return proto.CompactTextString(m) }
func (*ResendConfirmationEmailReq_Inner) ProtoMessage()    {}
func (*ResendConfirmationEmailReq_Inner) Descriptor() ([]byte, []int) {
	return fileDescriptorAccount, []int{3, 0}
}

type SignOutReq struct {
}

func (m *SignOutReq) Reset()                    { *m = SignOutReq{} }
func (m *SignOutReq) String() string            { return proto.CompactTextString(m) }
func (*SignOutReq) ProtoMessage()               {}
func (*SignOutReq) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{4} }

type ConfirmEmailReq struct {
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *ConfirmEmailReq) Reset()                    { *m = ConfirmEmailReq{} }
func (m *ConfirmEmailReq) String() string            { return proto.CompactTextString(m) }
func (*ConfirmEmailReq) ProtoMessage()               {}
func (*ConfirmEmailReq) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{5} }

func (m *ConfirmEmailReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ForgotPasswordReq struct {
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (m *ForgotPasswordReq) Reset()                    { *m = ForgotPasswordReq{} }
func (m *ForgotPasswordReq) String() string            { return proto.CompactTextString(m) }
func (*ForgotPasswordReq) ProtoMessage()               {}
func (*ForgotPasswordReq) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{6} }

func (m *ForgotPasswordReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type CheckResetPasswordTokenReq struct {
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *CheckResetPasswordTokenReq) Reset()         { *m = CheckResetPasswordTokenReq{} }
func (m *CheckResetPasswordTokenReq) String() string { return proto.CompactTextString(m) }
func (*CheckResetPasswordTokenReq) ProtoMessage()    {}
func (*CheckResetPasswordTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptorAccount, []int{7}
}

func (m *CheckResetPasswordTokenReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RequestChangeEmailReq struct {
	NewEmail string `protobuf:"bytes,1,opt,name=new_email,json=newEmail,proto3" json:"new_email,omitempty"`
}

func (m *RequestChangeEmailReq) Reset()                    { *m = RequestChangeEmailReq{} }
func (m *RequestChangeEmailReq) String() string            { return proto.CompactTextString(m) }
func (*RequestChangeEmailReq) ProtoMessage()               {}
func (*RequestChangeEmailReq) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{8} }

func (m *RequestChangeEmailReq) GetNewEmail() string {
	if m != nil {
		return m.NewEmail
	}
	return ""
}

type ChangeEmailReq struct {
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *ChangeEmailReq) Reset()                    { *m = ChangeEmailReq{} }
func (m *ChangeEmailReq) String() string            { return proto.CompactTextString(m) }
func (*ChangeEmailReq) ProtoMessage()               {}
func (*ChangeEmailReq) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{9} }

func (m *ChangeEmailReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ResetPasswordReq struct {
	Token    string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Password []string `protobuf:"bytes,2,rep,name=password" json:"password,omitempty"`
}

func (m *ResetPasswordReq) Reset()                    { *m = ResetPasswordReq{} }
func (m *ResetPasswordReq) String() string            { return proto.CompactTextString(m) }
func (*ResetPasswordReq) ProtoMessage()               {}
func (*ResetPasswordReq) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{10} }

func (m *ResetPasswordReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ResetPasswordReq) GetPassword() []string {
	if m != nil {
		return m.Password
	}
	return nil
}

type ChangePasswordReq struct {
	OldPassword string `protobuf:"bytes,1,opt,name=old_password,json=oldPassword,proto3" json:"old_password,omitempty"`
	NewPassword string `protobuf:"bytes,2,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
}

func (m *ChangePasswordReq) Reset()                    { *m = ChangePasswordReq{} }
func (m *ChangePasswordReq) String() string            { return proto.CompactTextString(m) }
func (*ChangePasswordReq) ProtoMessage()               {}
func (*ChangePasswordReq) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{11} }

func (m *ChangePasswordReq) GetOldPassword() string {
	if m != nil {
		return m.OldPassword
	}
	return ""
}

func (m *ChangePasswordReq) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

type RequestDeleteAccountReq struct {
}

func (m *RequestDeleteAccountReq) Reset()                    { *m = RequestDeleteAccountReq{} }
func (m *RequestDeleteAccountReq) String() string            { return proto.CompactTextString(m) }
func (*RequestDeleteAccountReq) ProtoMessage()               {}
func (*RequestDeleteAccountReq) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{12} }

type DeleteAccountReq struct {
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *DeleteAccountReq) Reset()                    { *m = DeleteAccountReq{} }
func (m *DeleteAccountReq) String() string            { return proto.CompactTextString(m) }
func (*DeleteAccountReq) ProtoMessage()               {}
func (*DeleteAccountReq) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{13} }

func (m *DeleteAccountReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// response
type SignInRes struct {
	Account *Account `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
}

func (m *SignInRes) Reset()                    { *m = SignInRes{} }
func (m *SignInRes) String() string            { return proto.CompactTextString(m) }
func (*SignInRes) ProtoMessage()               {}
func (*SignInRes) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{14} }

func (m *SignInRes) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type SignUpWithEmailRes struct {
	Account *Account `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
}

func (m *SignUpWithEmailRes) Reset()                    { *m = SignUpWithEmailRes{} }
func (m *SignUpWithEmailRes) String() string            { return proto.CompactTextString(m) }
func (*SignUpWithEmailRes) ProtoMessage()               {}
func (*SignUpWithEmailRes) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{15} }

func (m *SignUpWithEmailRes) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type ConfirmEmailRes struct {
}

func (m *ConfirmEmailRes) Reset()                    { *m = ConfirmEmailRes{} }
func (m *ConfirmEmailRes) String() string            { return proto.CompactTextString(m) }
func (*ConfirmEmailRes) ProtoMessage()               {}
func (*ConfirmEmailRes) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{16} }

type GetCurrentAccountRes struct {
	Account []*Account `protobuf:"bytes,1,rep,name=account" json:"account,omitempty"`
}

func (m *GetCurrentAccountRes) Reset()                    { *m = GetCurrentAccountRes{} }
func (m *GetCurrentAccountRes) String() string            { return proto.CompactTextString(m) }
func (*GetCurrentAccountRes) ProtoMessage()               {}
func (*GetCurrentAccountRes) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{17} }

func (m *GetCurrentAccountRes) GetAccount() []*Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type CheckResetPasswordTokenRes struct {
}

func (m *CheckResetPasswordTokenRes) Reset()         { *m = CheckResetPasswordTokenRes{} }
func (m *CheckResetPasswordTokenRes) String() string { return proto.CompactTextString(m) }
func (*CheckResetPasswordTokenRes) ProtoMessage()    {}
func (*CheckResetPasswordTokenRes) Descriptor() ([]byte, []int) {
	return fileDescriptorAccount, []int{18}
}

type SignOutRes struct {
}

func (m *SignOutRes) Reset()                    { *m = SignOutRes{} }
func (m *SignOutRes) String() string            { return proto.CompactTextString(m) }
func (*SignOutRes) ProtoMessage()               {}
func (*SignOutRes) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{19} }

type ResendConfirmationEmailRes struct {
}

func (m *ResendConfirmationEmailRes) Reset()         { *m = ResendConfirmationEmailRes{} }
func (m *ResendConfirmationEmailRes) String() string { return proto.CompactTextString(m) }
func (*ResendConfirmationEmailRes) ProtoMessage()    {}
func (*ResendConfirmationEmailRes) Descriptor() ([]byte, []int) {
	return fileDescriptorAccount, []int{20}
}

type ForgotPasswordRes struct {
}

func (m *ForgotPasswordRes) Reset()                    { *m = ForgotPasswordRes{} }
func (m *ForgotPasswordRes) String() string            { return proto.CompactTextString(m) }
func (*ForgotPasswordRes) ProtoMessage()               {}
func (*ForgotPasswordRes) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{21} }

type ResetPasswordRes struct {
}

func (m *ResetPasswordRes) Reset()                    { *m = ResetPasswordRes{} }
func (m *ResetPasswordRes) String() string            { return proto.CompactTextString(m) }
func (*ResetPasswordRes) ProtoMessage()               {}
func (*ResetPasswordRes) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{22} }

type ChangePasswordRes struct {
}

func (m *ChangePasswordRes) Reset()                    { *m = ChangePasswordRes{} }
func (m *ChangePasswordRes) String() string            { return proto.CompactTextString(m) }
func (*ChangePasswordRes) ProtoMessage()               {}
func (*ChangePasswordRes) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{23} }

type RequestChangeEmailRes struct {
}

func (m *RequestChangeEmailRes) Reset()                    { *m = RequestChangeEmailRes{} }
func (m *RequestChangeEmailRes) String() string            { return proto.CompactTextString(m) }
func (*RequestChangeEmailRes) ProtoMessage()               {}
func (*RequestChangeEmailRes) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{24} }

type ChangeEmailRes struct {
	Account *Account `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
}

func (m *ChangeEmailRes) Reset()                    { *m = ChangeEmailRes{} }
func (m *ChangeEmailRes) String() string            { return proto.CompactTextString(m) }
func (*ChangeEmailRes) ProtoMessage()               {}
func (*ChangeEmailRes) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{25} }

func (m *ChangeEmailRes) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type RequestDeleteAccountRes struct {
}

func (m *RequestDeleteAccountRes) Reset()                    { *m = RequestDeleteAccountRes{} }
func (m *RequestDeleteAccountRes) String() string            { return proto.CompactTextString(m) }
func (*RequestDeleteAccountRes) ProtoMessage()               {}
func (*RequestDeleteAccountRes) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{26} }

type DeleteAccountRes struct {
}

func (m *DeleteAccountRes) Reset()                    { *m = DeleteAccountRes{} }
func (m *DeleteAccountRes) String() string            { return proto.CompactTextString(m) }
func (*DeleteAccountRes) ProtoMessage()               {}
func (*DeleteAccountRes) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{27} }

// common
type Account struct {
	AccountId      string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Email          string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	EmailConfirmed bool   `protobuf:"varint,3,opt,name=email_confirmed,json=emailConfirmed,proto3" json:"email_confirmed,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{28} }

func (m *Account) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Account) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Account) GetEmailConfirmed() bool {
	if m != nil {
		return m.EmailConfirmed
	}
	return false
}

// events
type SignedUpWithEmail struct {
	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password []byte `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (m *SignedUpWithEmail) Reset()                    { *m = SignedUpWithEmail{} }
func (m *SignedUpWithEmail) String() string            { return proto.CompactTextString(m) }
func (*SignedUpWithEmail) ProtoMessage()               {}
func (*SignedUpWithEmail) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{29} }

func (m *SignedUpWithEmail) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignedUpWithEmail) GetPassword() []byte {
	if m != nil {
		return m.Password
	}
	return nil
}

type EmailConfirmed struct {
}

func (m *EmailConfirmed) Reset()                    { *m = EmailConfirmed{} }
func (m *EmailConfirmed) String() string            { return proto.CompactTextString(m) }
func (*EmailConfirmed) ProtoMessage()               {}
func (*EmailConfirmed) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{30} }

type PasswordReset struct {
	Password []byte `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
}

func (m *PasswordReset) Reset()                    { *m = PasswordReset{} }
func (m *PasswordReset) String() string            { return proto.CompactTextString(m) }
func (*PasswordReset) ProtoMessage()               {}
func (*PasswordReset) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{31} }

func (m *PasswordReset) GetPassword() []byte {
	if m != nil {
		return m.Password
	}
	return nil
}

type PasswordChanged struct {
	Password []byte `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
}

func (m *PasswordChanged) Reset()                    { *m = PasswordChanged{} }
func (m *PasswordChanged) String() string            { return proto.CompactTextString(m) }
func (*PasswordChanged) ProtoMessage()               {}
func (*PasswordChanged) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{32} }

func (m *PasswordChanged) GetPassword() []byte {
	if m != nil {
		return m.Password
	}
	return nil
}

type EmailChanged struct {
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (m *EmailChanged) Reset()                    { *m = EmailChanged{} }
func (m *EmailChanged) String() string            { return proto.CompactTextString(m) }
func (*EmailChanged) ProtoMessage()               {}
func (*EmailChanged) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{33} }

func (m *EmailChanged) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type AccountDeleted struct {
}

func (m *AccountDeleted) Reset()                    { *m = AccountDeleted{} }
func (m *AccountDeleted) String() string            { return proto.CompactTextString(m) }
func (*AccountDeleted) ProtoMessage()               {}
func (*AccountDeleted) Descriptor() ([]byte, []int) { return fileDescriptorAccount, []int{34} }

func init() {
	proto.RegisterType((*SignInReq)(nil), "account.SignInReq")
	proto.RegisterType((*GetCurrentAccountReq)(nil), "account.GetCurrentAccountReq")
	proto.RegisterType((*SignUpWithEmailReq)(nil), "account.SignUpWithEmailReq")
	proto.RegisterType((*ResendConfirmationEmailReq)(nil), "account.ResendConfirmationEmailReq")
	proto.RegisterType((*ResendConfirmationEmailReq_Inner)(nil), "account.ResendConfirmationEmailReq.Inner")
	proto.RegisterType((*SignOutReq)(nil), "account.SignOutReq")
	proto.RegisterType((*ConfirmEmailReq)(nil), "account.ConfirmEmailReq")
	proto.RegisterType((*ForgotPasswordReq)(nil), "account.ForgotPasswordReq")
	proto.RegisterType((*CheckResetPasswordTokenReq)(nil), "account.CheckResetPasswordTokenReq")
	proto.RegisterType((*RequestChangeEmailReq)(nil), "account.RequestChangeEmailReq")
	proto.RegisterType((*ChangeEmailReq)(nil), "account.ChangeEmailReq")
	proto.RegisterType((*ResetPasswordReq)(nil), "account.ResetPasswordReq")
	proto.RegisterType((*ChangePasswordReq)(nil), "account.ChangePasswordReq")
	proto.RegisterType((*RequestDeleteAccountReq)(nil), "account.RequestDeleteAccountReq")
	proto.RegisterType((*DeleteAccountReq)(nil), "account.DeleteAccountReq")
	proto.RegisterType((*SignInRes)(nil), "account.SignInRes")
	proto.RegisterType((*SignUpWithEmailRes)(nil), "account.SignUpWithEmailRes")
	proto.RegisterType((*ConfirmEmailRes)(nil), "account.ConfirmEmailRes")
	proto.RegisterType((*GetCurrentAccountRes)(nil), "account.GetCurrentAccountRes")
	proto.RegisterType((*CheckResetPasswordTokenRes)(nil), "account.CheckResetPasswordTokenRes")
	proto.RegisterType((*SignOutRes)(nil), "account.SignOutRes")
	proto.RegisterType((*ResendConfirmationEmailRes)(nil), "account.ResendConfirmationEmailRes")
	proto.RegisterType((*ForgotPasswordRes)(nil), "account.ForgotPasswordRes")
	proto.RegisterType((*ResetPasswordRes)(nil), "account.ResetPasswordRes")
	proto.RegisterType((*ChangePasswordRes)(nil), "account.ChangePasswordRes")
	proto.RegisterType((*RequestChangeEmailRes)(nil), "account.RequestChangeEmailRes")
	proto.RegisterType((*ChangeEmailRes)(nil), "account.ChangeEmailRes")
	proto.RegisterType((*RequestDeleteAccountRes)(nil), "account.RequestDeleteAccountRes")
	proto.RegisterType((*DeleteAccountRes)(nil), "account.DeleteAccountRes")
	proto.RegisterType((*Account)(nil), "account.Account")
	proto.RegisterType((*SignedUpWithEmail)(nil), "account.SignedUpWithEmail")
	proto.RegisterType((*EmailConfirmed)(nil), "account.EmailConfirmed")
	proto.RegisterType((*PasswordReset)(nil), "account.PasswordReset")
	proto.RegisterType((*PasswordChanged)(nil), "account.PasswordChanged")
	proto.RegisterType((*EmailChanged)(nil), "account.EmailChanged")
	proto.RegisterType((*AccountDeleted)(nil), "account.AccountDeleted")
	proto.RegisterEnum("account.Roles", Roles_name, Roles_value)
	proto.RegisterEnum("account.ResendConfirmationEmailReq_Data", ResendConfirmationEmailReq_Data_name, ResendConfirmationEmailReq_Data_value)
	proto.RegisterEnum("account.ResendConfirmationEmailReq_Inner_InnerEnum", ResendConfirmationEmailReq_Inner_InnerEnum_name, ResendConfirmationEmailReq_Inner_InnerEnum_value)
}

func init() { proto.RegisterFile("account.proto", fileDescriptorAccount) }

var fileDescriptorAccount = []byte{
	// 996 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0x7d, 0x6f, 0xdb, 0x44,
	0x18, 0xc7, 0x69, 0xd3, 0x36, 0x4f, 0xd3, 0x34, 0x39, 0x0a, 0xed, 0xcc, 0x4a, 0xcb, 0xb1, 0x8d,
	0x6a, 0xc5, 0xc9, 0xc8, 0x78, 0x11, 0x68, 0x12, 0x5b, 0x93, 0x50, 0xaa, 0x4a, 0x05, 0x9c, 0x21,
	0x2a, 0xd8, 0xc8, 0xbc, 0xf8, 0x48, 0xad, 0x3a, 0xbe, 0xc6, 0xe7, 0x34, 0x08, 0x84, 0xc4, 0xbf,
	0xf0, 0x05, 0xf8, 0x3a, 0x7c, 0x93, 0x4a, 0xe5, 0x8b, 0xa0, 0xbb, 0xb3, 0x5d, 0xbf, 0xc4, 0x6e,
	0x26, 0x75, 0xff, 0x54, 0xbe, 0x7b, 0x7e, 0xcf, 0xef, 0x9e, 0xfb, 0xdd, 0xf3, 0x92, 0xc2, 0x8a,
	0xd1, 0xef, 0xd3, 0xb1, 0xe3, 0xd5, 0xcf, 0x5c, 0xea, 0x51, 0xb4, 0xe8, 0x2f, 0xd5, 0xfa, 0xc0,
	0xf2, 0x4e, 0xc6, 0x2f, 0xeb, 0x7d, 0x3a, 0x6c, 0x98, 0x86, 0x63, 0x11, 0xfb, 0xdc, 0x36, 0xcc,
	0x3e, 0x6d, 0x0c, 0xa8, 0x26, 0xa0, 0xda, 0x60, 0x64, 0x37, 0x06, 0x23, 0x5b, 0x3a, 0xaa, 0x9f,
	0x46, 0xf0, 0xc3, 0x89, 0xe5, 0x9d, 0xd2, 0xc9, 0x15, 0xf4, 0xdc, 0xb0, 0x2d, 0xd3, 0xf0, 0xa8,
	0xcb, 0x1a, 0xe1, 0xa7, 0xf4, 0xc3, 0x7f, 0x2a, 0x50, 0xea, 0x5a, 0x03, 0xe7, 0xc0, 0xd1, 0xc9,
	0x08, 0x1d, 0x42, 0x91, 0x0c, 0x0d, 0xcb, 0xde, 0x50, 0xb6, 0x95, 0x9d, 0xd2, 0xde, 0x27, 0x97,
	0x17, 0x5b, 0x1f, 0x81, 0xf6, 0xf3, 0x4f, 0x86, 0xf6, 0xdb, 0x13, 0xed, 0xc7, 0x07, 0xda, 0xe7,
	0xf5, 0xde, 0xdd, 0xdd, 0x67, 0xda, 0xf3, 0xdd, 0xc7, 0x7c, 0x8b, 0xaf, 0xf9, 0xe2, 0x59, 0x9d,
	0xaf, 0x9e, 0xff, 0xde, 0xfc, 0xf0, 0xe3, 0x3f, 0xee, 0x1c, 0x2b, 0xba, 0xe4, 0x40, 0xf7, 0x60,
	0xe9, 0xcc, 0x60, 0x6c, 0x42, 0x5d, 0x73, 0xa3, 0x20, 0xf8, 0xe0, 0xf2, 0x62, 0x6b, 0xe1, 0x58,
	0x39, 0x2b, 0xfe, 0x6a, 0xea, 0xa1, 0x0d, 0xbf, 0x0d, 0x6b, 0xfb, 0xc4, 0x6b, 0x8d, 0x5d, 0x97,
	0x38, 0xde, 0x13, 0x79, 0x7f, 0x9d, 0x8c, 0xf0, 0x5f, 0x0a, 0x20, 0x1e, 0xda, 0xf7, 0x67, 0x3f,
	0x58, 0xde, 0x49, 0x87, 0x73, 0xbe, 0xd6, 0x18, 0xe7, 0x73, 0x62, 0xfc, 0xa7, 0x00, 0xaa, 0x4e,
	0x18, 0x71, 0xcc, 0x16, 0x75, 0x7e, 0xb1, 0xdc, 0xa1, 0xe1, 0x59, 0xd4, 0x09, 0x63, 0x3a, 0x80,
	0xa2, 0xe5, 0x38, 0xc4, 0x15, 0x1c, 0x95, 0xe6, 0xc3, 0x7a, 0xf0, 0xaa, 0xd9, 0x3e, 0xf5, 0x03,
	0xee, 0x20, 0xff, 0x76, 0x9c, 0xf1, 0x50, 0x97, 0x0c, 0xe8, 0x11, 0xcc, 0x9b, 0x86, 0x67, 0x08,
	0xc5, 0x2a, 0xcd, 0x9d, 0x59, 0x98, 0xda, 0x86, 0x67, 0xe8, 0xc2, 0x0b, 0xdd, 0x05, 0xf0, 0x1d,
	0x7a, 0x96, 0xe9, 0x2b, 0xb4, 0x70, 0x79, 0xb1, 0x55, 0x38, 0x56, 0xf4, 0x92, 0x6f, 0x39, 0x30,
	0xd5, 0x7b, 0x50, 0x14, 0x07, 0xe3, 0x4d, 0x28, 0x85, 0x11, 0xa0, 0x45, 0x98, 0xeb, 0x76, 0xbb,
	0xd5, 0x37, 0xf8, 0x47, 0xbb, 0xdd, 0xae, 0x2a, 0xf8, 0x3d, 0x98, 0xe7, 0xe4, 0xc1, 0x86, 0xb0,
	0x1c, 0x1d, 0x1d, 0x55, 0x15, 0xfe, 0xb1, 0xbf, 0xbf, 0x5f, 0x2d, 0xe0, 0x32, 0x00, 0x7f, 0xa4,
	0x6f, 0xc6, 0xe2, 0xcd, 0x1a, 0xb0, 0xea, 0x87, 0x18, 0x6a, 0x73, 0x1b, 0x8a, 0x1e, 0x3d, 0x25,
	0x4e, 0x22, 0x1a, 0xb9, 0x89, 0x5f, 0x40, 0xed, 0x2b, 0xea, 0x0e, 0xa8, 0xf7, 0xad, 0x2f, 0xf5,
	0x4d, 0x3f, 0x31, 0xfe, 0x02, 0xd4, 0xd6, 0x09, 0xe9, 0x9f, 0x72, 0x01, 0xc3, 0x53, 0x9e, 0xf2,
	0xc3, 0xaf, 0x8f, 0xee, 0x14, 0xde, 0xd2, 0xc9, 0x68, 0x4c, 0x98, 0xd7, 0x3a, 0x31, 0x9c, 0x01,
	0x09, 0x2f, 0xa5, 0x43, 0xc9, 0x21, 0x93, 0xde, 0x0d, 0x44, 0xb9, 0xe4, 0x90, 0x89, 0xa0, 0xc5,
	0x75, 0xa8, 0x24, 0x4e, 0xc9, 0x0f, 0xee, 0x18, 0xaa, 0xb1, 0x3b, 0x5d, 0xeb, 0x91, 0xa8, 0xc8,
	0xb9, 0xcc, 0x6c, 0x1f, 0x41, 0x4d, 0x46, 0x12, 0xa5, 0xd6, 0xa0, 0x4c, 0x6d, 0xb3, 0x17, 0x12,
	0x28, 0xa9, 0x72, 0x59, 0xa6, 0xb6, 0x19, 0x78, 0x70, 0x38, 0x57, 0x28, 0xa7, 0x03, 0x2c, 0x3b,
	0x64, 0x12, 0xc0, 0xf1, 0x2d, 0x58, 0xf7, 0x95, 0x6e, 0x13, 0x9b, 0x78, 0x24, 0xd2, 0x07, 0x76,
	0xa0, 0x9a, 0xdc, 0x43, 0x6b, 0xb1, 0x7b, 0x06, 0x8a, 0x7c, 0x76, 0xd5, 0xcb, 0x18, 0xba, 0x0f,
	0x41, 0x33, 0x15, 0xa0, 0xe5, 0x66, 0x35, 0xac, 0xa5, 0x80, 0x28, 0x00, 0xe0, 0xc7, 0x53, 0x3a,
	0xcd, 0xab, 0x31, 0xd4, 0x92, 0x89, 0xcf, 0xf0, 0xde, 0xd4, 0xbe, 0x96, 0xa0, 0x9d, 0xcb, 0xa7,
	0xbd, 0x9d, 0x93, 0xbc, 0x2c, 0x56, 0x7b, 0x8c, 0x63, 0x33, 0x9b, 0x04, 0xc3, 0x6f, 0xa6, 0x0b,
	0x8d, 0x61, 0x94, 0x4a, 0x21, 0x01, 0x4c, 0x3e, 0x3e, 0xc3, 0xeb, 0xd3, 0x0b, 0x81, 0xe1, 0x47,
	0x89, 0xa4, 0x7d, 0x35, 0xd5, 0x32, 0x5f, 0x5d, 0x84, 0x96, 0xda, 0x1b, 0xc0, 0xa2, 0xbf, 0x42,
	0x9b, 0xe9, 0x46, 0x17, 0x69, 0x70, 0x3c, 0x3f, 0x64, 0x6d, 0x16, 0x64, 0x7e, 0xc8, 0x6e, 0xff,
	0x01, 0xac, 0x8a, 0x8f, 0x5e, 0x5f, 0x2a, 0x44, 0xcc, 0x8d, 0xb9, 0x6d, 0x65, 0x67, 0x49, 0xaf,
	0x88, 0xed, 0x56, 0xb0, 0x8b, 0x3b, 0x50, 0xe3, 0xc2, 0x12, 0x33, 0x92, 0x11, 0x57, 0x9c, 0x4a,
	0x94, 0x53, 0x4d, 0x4c, 0x90, 0x72, 0xa4, 0x8e, 0xaa, 0x50, 0xe9, 0xc4, 0x89, 0x77, 0x61, 0x25,
	0x22, 0x2b, 0xf1, 0x62, 0xee, 0x4a, 0xc2, 0x5d, 0x83, 0xd5, 0x00, 0x2c, 0x35, 0x36, 0x73, 0xe1,
	0x77, 0xa0, 0x2c, 0x4f, 0xf3, 0xb1, 0x53, 0xe3, 0xe5, 0x31, 0xf9, 0x1a, 0x4a, 0x79, 0xcd, 0xfb,
	0x18, 0x8a, 0x3a, 0xb5, 0x09, 0xcb, 0xe9, 0xf2, 0xcd, 0xbf, 0x97, 0x60, 0xb1, 0x4b, 0xdc, 0x73,
	0xab, 0x4f, 0xd0, 0x03, 0x58, 0x90, 0x55, 0x86, 0x50, 0xf8, 0xb2, 0xe1, 0x4f, 0x08, 0x35, 0xbd,
	0xc7, 0xd0, 0x77, 0x50, 0x4b, 0x55, 0x02, 0xda, 0x0c, 0x81, 0xd3, 0xa6, 0xbf, 0x9a, 0x6b, 0x66,
	0xe8, 0x10, 0x56, 0x13, 0x15, 0x8b, 0xde, 0x89, 0x9d, 0x1c, 0xff, 0xd5, 0xa0, 0xe6, 0x18, 0x19,
	0xea, 0xf3, 0x34, 0x9c, 0x5a, 0x39, 0xe8, 0xfd, 0x19, 0x06, 0xb0, 0x3a, 0x03, 0x88, 0xa1, 0x3d,
	0x28, 0x47, 0x3b, 0x04, 0xda, 0x08, 0x9d, 0x12, 0x13, 0x53, 0xcd, 0xb2, 0x30, 0xf4, 0x35, 0x54,
	0xe2, 0x45, 0x8c, 0xd4, 0x10, 0x9b, 0x1a, 0xa3, 0x6a, 0xb6, 0x4d, 0x5c, 0x39, 0xa3, 0xb1, 0x44,
	0xae, 0x9c, 0x3d, 0x37, 0xd5, 0x19, 0x40, 0x0c, 0x75, 0x60, 0x25, 0x66, 0x40, 0xb7, 0x62, 0x42,
	0xc5, 0x82, 0xcd, 0x34, 0x89, 0x5b, 0xc7, 0x3b, 0x52, 0xe4, 0xd6, 0xa9, 0x39, 0xa5, 0x66, 0xdb,
	0x18, 0x7a, 0x0a, 0x28, 0xdd, 0xc6, 0xd0, 0xbb, 0x91, 0xa3, 0xa7, 0x0c, 0x7b, 0x35, 0xdf, 0xce,
	0xd0, 0x97, 0xb0, 0x1c, 0xa5, 0x5b, 0x4f, 0x04, 0x10, 0xf2, 0x64, 0x18, 0x18, 0x7a, 0x01, 0x6b,
	0xd3, 0xda, 0x20, 0xda, 0x4e, 0x1e, 0x9c, 0x9c, 0x83, 0xea, 0x75, 0x08, 0x86, 0xe7, 0xff, 0xfd,
	0x6f, 0xab, 0x80, 0x0e, 0x61, 0x25, 0x4e, 0x7d, 0x25, 0x77, 0x8a, 0x33, 0xd3, 0x24, 0xc9, 0x94,
	0x97, 0x0b, 0xe2, 0x5f, 0x87, 0x87, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xfb, 0xb2, 0x91,
	0xbc, 0x0c, 0x00, 0x00,
}