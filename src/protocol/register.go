package protocol

func init() {
	c2s_protobufs[int32(Default_C2S_Login_Type)] = &C2S_Login{}
	s2c_protobufs[int32(Default_S2C_Login_Type)] = &S2C_Login{}
	c2s_protobufs[int32(Default_C2S_RandName_Type)] = &C2S_RandName{}
	s2c_protobufs[int32(Default_S2C_RandName_Type)] = &S2C_RandName{}
	c2s_protobufs[int32(Default_C2S_CreateRole_Type)] = &C2S_CreateRole{}
	s2c_protobufs[int32(Default_S2C_CreateRole_Type)] = &S2C_CreateRole{}
	c2s_protobufs[int32(Default_C2S_LoadRoleInfo_Type)] = &C2S_LoadRoleInfo{}
	s2c_protobufs[int32(Default_S2C_LoadRoleInfo_Type)] = &S2C_LoadRoleInfo{}
	s2c_protobufs[int32(Default_S2C_ServerTime_Type)] = &S2C_ServerTime{}
}
