package state

type Action uint8

const (
	ActionNop         Action = 0
	ActionClear       Action = 1
	ActionCollect     Action = 2
	ActionCSIDispatch Action = 3
	ActionESCDispatch Action = 4
	ActionExecute     Action = 5
	ActionHook        Action = 6
	ActionIgnore      Action = 7
	ActionOSCEnd      Action = 8
	ActionOSCPut      Action = 9
	ActionOSCStart    Action = 10
	ActionParam       Action = 11
	ActionPrint       Action = 12
	ActionPut         Action = 13
	ActionUnhook      Action = 14
	ActionBeginUTF8   Action = 15
)
