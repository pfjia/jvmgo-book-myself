package instructions

import (
	"awesomeProject/ch05/instructions/base"
	"awesomeProject/ch05/instructions/comparisons"
	"awesomeProject/ch05/instructions/constants"
	"awesomeProject/ch05/instructions/control"
	"awesomeProject/ch05/instructions/conversions"
	"awesomeProject/ch05/instructions/loads"
	"awesomeProject/ch05/instructions/math"
	"awesomeProject/ch05/instructions/stack"
	"awesomeProject/ch05/instructions/stores"
	"fmt"
)

func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00:
		return &constants.NOP{}
	case 0x01:
		return &constants.ACONST_NULL{}
	case 0x02:
		return &constants.ICONST_M1{}
	case 0x03:
		return &constants.ICONST_0{}
	case 0x04:
		return &constants.ICONST_1{}
	case 0x05:
		return &constants.ICONST_2{}
	case 0x06:
		return &constants.ICONST_3{}
	case 0x07:
		return &constants.ICONST_4{}
	case 0x08:
		return &constants.ICONST_5{}
	case 0x09:
		return &constants.LCONST_0{}
	case 0x0a:
		return &constants.LCONST_1{}
	case 0x0b:
		return &constants.FCONST_0{}
	case 0x0c:
		return &constants.FCONST_1{}
	case 0x0d:
		return &constants.FCONST_2{}
	case 0x0e:
		return &constants.DCONST_0{}
	case 0x0f:
		return &constants.DCONST_1{}
	case 0x10:
		return &constants.BIPUSH{}
	case 0x11:
		return &constants.SIPUSH{}
	//case 0x12:
	//	return &constants.LDC{}
	//case 0x13:
	//	return &constants.LDC_W{}
	//case 0x14:
	//	return &constants.LDC2_W{}
	case 0x15:
		return &loads.ILOAD{}
	case 0x16:
		return &loads.LLOAD{}
	case 0x17:
		return &loads.FLOAD{}
	case 0x18:
		return &loads.DLOAD{}
	case 0x19:
		return &loads.ALOAD{}
	case 0x1a:
		return &loads.ILOAD_0{}
	case 0x1b:
		return &loads.ILOAD_1{}
	case 0x1c:
		return &loads.ILOAD_2{}
	case 0x1d:
		return &loads.ILOAD_3{}
	case 0x1e:
		return &loads.LLOAD_0{}
	case 0x1f:
		return &loads.LLOAD_1{}
	case 0x20:
		return &loads.LLOAD_2{}
	case 0x21:
		return &loads.LLOAD_3{}
	case 0x22:
		return &loads.FLOAD_0{}
	case 0x23:
		return &loads.FLOAD_1{}
	case 0x24:
		return &loads.FLOAD_2{}
	case 0x25:
		return &loads.FLOAD_3{}
	case 0x26:
		return &loads.DLOAD_0{}
	case 0x27:
		return &loads.DLOAD_1{}
	case 0x28:
		return &loads.DLOAD_2{}
	case 0x29:
		return &loads.DLOAD_3{}
	case 0x2a:
		return &loads.ALOAD_0{}
	case 0x2b:
		return &loads.ALOAD_1{}
	case 0x2c:
		return &loads.ALOAD_2{}
	case 0x2d:
		return &loads.ALOAD_3{}
	//case 0x2e:
	//	return &loads.IALOAD{}
	//case 0x2f:
	//	return &loads.LALOAD{}
	//case 0x30:
	//	return &loads.FALOAD{}
	//case 0x31:
	//	return &loads.DALOAD{}
	//case 0x32:
	//	return &loads.AALOAD{}
	//case 0x33:
	//	return &loads.BALOAD{}
	//case 0x34:
	//	return &loads.CALOAD{}
	//case 0x35:
	//	return &loads.SALOAD{}
	case 0x36:
		return &stores.ISTORE{}
	case 0x37:
		return &stores.LSTORE{}
	case 0x38:
		return &stores.FSTORE{}
	case 0x39:
		return &stores.DSTORE{}
	case 0x3a:
		return &stores.ASTORE{}
	case 0x3b:
		return &stores.ISTORE_0{}
	case 0x3c:
		return &stores.ISTORE_1{}
	case 0x3d:
		return &stores.ISTORE_2{}
	case 0x3e:
		return &stores.ISTORE_3{}
	case 0x3f:
		return &stores.LSTORE_0{}
	case 0x40:
		return &stores.LSTORE_1{}
	case 0x41:
		return &stores.LSTORE_2{}
	case 0x42:
		return &stores.LSTORE_3{}
	case 0x43:
		return &stores.FSTORE_0{}
	case 0x44:
		return &stores.FSTORE_1{}
	case 0x45:
		return &stores.FSTORE_2{}
	case 0x46:
		return &stores.FSTORE_3{}
	case 0x47:
		return &stores.DSTORE_0{}
	case 0x48:
		return &stores.DSTORE_1{}
	case 0x49:
		return &stores.DSTORE_2{}
	case 0x4a:
		return &stores.DSTORE_3{}
	case 0x4b:
		return &stores.ASTORE_0{}
	case 0x4c:
		return &stores.ASTORE_1{}
	case 0x4d:
		return &stores.ASTORE_2{}
	case 0x4e:
		return &stores.ASTORE_3{}
	//case 0x4f:
	//	return &stores.IASTORE{}
	//case 0x50:
	//	return &stores.LASTORE{}
	//case 0x51:
	//	return &stores.FASTORE{}
	//case 0x52:
	//	return &stores.DASTORE{}
	//case 0x53:
	//	return &stores.AASTORE{}
	//case 0x54:
	//	return &stores.BASTORE{}
	//case 0x55:
	//	return &stores.CASTORE{}
	//case 0x56:
	//	return &stores.SASTORE{}
	case 0x57:
		return &stack.POP{}
	case 0x58:
		return &stack.POP2{}
	case 0x59:
		return &stack.DUP{}
	case 0x5a:
		return &stack.DUP_X1{}
	case 0x5b:
		return &stack.DUP_X2{}
	case 0x5c:
		return &stack.DUP2{}
	case 0x5d:
		return &stack.DUP2_X1{}
	case 0x5e:
		return &stack.DUP2_X2{}
	case 0x5f:
		return &stack.SWAP{}
	case 0x60:
		return &math.IADD{}
	case 0x61:
		return &math.LADD{}
	case 0x62:
		return &math.FADD{}
	case 0x63:
		return &math.DADD{}
	case 0x64:
		return &math.ISUB{}
	case 0x65:
		return &math.LSUB{}
	case 0x66:
		return &math.FSUB{}
	case 0x67:
		return &math.DSUB{}
	case 0x68:
		return &math.IMUL{}
	case 0x69:
		return &math.LMUL{}
	case 0x6a:
		return &math.FMUL{}
	case 0x6b:
		return &math.DMUL{}
	case 0x6c:
		return &math.IDIV{}
	case 0x6d:
		return &math.LDIV{}
	case 0x6e:
		return &math.FDIV{}
	case 0x6f:
		return &math.DDIV{}
	case 0x70:
		return &math.IREM{}
	case 0x71:
		return &math.LREM{}
	case 0x72:
		return &math.FREM{}
	case 0x73:
		return &math.DREM{}
	//case 0x74:
	//	return &math.INEG{}
	//case 0x75:
	//	return &math.LNEG{}
	//case 0x76:
	//	return &math.FNEG{}
	//case 0x77:
	//	return &math.DNEG{}
	case 0x78:
		return &math.ISHL{}
	case 0x79:
		return &math.LSHL{}
	case 0x7a:
		return &math.ISHR{}
	case 0x7b:
		return &math.LSHR{}
	case 0x7c:
		return &math.IUSHR{}
	case 0x7d:
		return &math.LUSHR{}
	case 0x7e:
		return &math.IAND{}
	case 0x7f:
		return &math.LAND{}
	//case 0x80:
	//	return &math.IOR{}
	//case 0x81:
	//	return &math.LOR{}
	//case 0x82:
	//	return &math.IXOR{}
	//case 0x83:
	//	return &math.LXOR{}
	case 0x84:
		return &math.IINC{}
	//case 0x85:
	//	return &conversions.I2L{}
	//case 0x86:
	//	return &conversions.I2F{}
	//case 0x87:
	//	return &conversions.I2D{}
	//case 0x88:
	//	return &conversions.L2I{}
	//case 0x89:
	//	return &conversions.L2F{}
	//case 0x8a:
	//	return &conversions.L2D{}
	//case 0x8b:
	//	return &conversions.F2I{}
	//case 0x8c:
	//	return &conversions.F2L{}
	//case 0x8d:
	//	return &conversions.F2D{}
	case 0x8e:
		return &conversions.D2I{}
	case 0x8f:
		return &conversions.D2L{}
	case 0x90:
		return &conversions.D2F{}
	//case 0x91:
	//	return &conversions.I2B{}
	//case 0x92:
	//	return &conversions.I2C{}
	//case 0x93:
	//	return &conversions.I2S{}
	case 0x94:
		return &comparisons.LCMP{}
	case 0x95:
		return &comparisons.FCMPL{}
	case 0x96:
		return &comparisons.FCMPG{}
	//case 0x97:
	//	return &comparisons.DCMPL{}
	//case 0x98:
	//	return &comparisons.DCMPG{}
	case 0x99:
		return &comparisons.IFEQ{}
	case 0x9a:
		return &comparisons.IFNE{}
	case 0x9b:
		return &comparisons.IFLT{}
	case 0x9c:
		return &comparisons.IFGE{}
	case 0x9d:
		return &comparisons.IFGT{}
	case 0x9e:
		return &comparisons.IFLE{}
	case 0x9f:
		return &comparisons.IF_ICMPEQ{}
	case 0xa0:
		return &comparisons.IF_ICMPNE{}
	case 0xa1:
		return &comparisons.IF_ICMPLT{}
	case 0xa2:
		return &comparisons.IF_ICMPGE{}
	case 0xa3:
		return &comparisons.IF_ICMPGT{}
	//case 0xa4:
	//	return &control.IF_ICMPLE{}
	//case 0xa5:
	//	return &control.IF_ACMPEQ{}
	//case 0xa6:
	//	return &control.IF_ACMPNE{}
	case 0xa7:
		return &control.GOTO{}
	//case 0xa8:
	//	return &control.JSR{}
	//case 0xa9:
	//	return &control.RET{}
	case 0xaa:
		return &control.TABLE_SWITCH{}
	case 0xab:
		return &control.LOOKUP_SWITCH{}
	//case 0xac:
	//	return &returns.IRETURN{}
	//case 0xad:
	//	return &returns.LRETURN{}
	//case 0xae:
	//	return &returns.FRETURN{}
	//case 0xaf:
	//	return &returns.DRETURN{}
	//case 0xb0:
	//	return &returns.ARETURN{}
	//case 0xb1:
	//	return &returns.RETURN{}
	//case 0xb2:
	//	return &references.GET_STATIC{}
	//case 0xb3:
	//	return &references.PUT_STATIC{}
	//case 0xb4:
	//	return &references.GET_FIELD{}
	//case 0xb5:
	//	return &references.PUT_FIELD{}
	//case 0xb6:
	//	return &references.INVOKE_VIRTUAL{}
	//case 0xb7:
	//	return &references.INVOKE_SPECIAL{}
	//case 0xb8:
	//	return &references.INVOKE_STATIC{}
	//case 0xb9:
	//	return &references.INVOKE_INTERFACE{}
	//case 0xba:
	//	return &references.INVOKE_DYNAMIC{}
	//case 0xbb:
	//	return &references.NEW{}
	//case 0xbc:
	//	return &references.NEW_ARRAY{}
	//case 0xbd:
	//	return &references.ANEW_ARRAY{}
	//case 0xbe:
	//	return &references.ARRAY_LENGTH{}
	//case 0xbf:
	//	return &references.ATHROW{}
	//case 0xc0:
	//	return &references.CHECK_CAST{}
	//case 0xc1:
	//	return &references.INSTANCE_OF{}
	//case 0xc2:
	//	return &references.MONITORENTER{}
	//case 0xc3:
	//	return &references.MONITOREXIT{}
	//case 0xc4:
	//	return &references.WIDE{}
	//case 0xc5:
	//	return &references.MULTIANEW_ARRAY{}
	//case 0xc6:
	//	return &references.IFNULL{}
	//case 0xc7:
	//	return &references.IFNONNULL{}
	//case 0xc8:
	//	return &references.GOTO_W{}
	//case 0xc9:
	//	return &references.JSR_W{}
	//case 0xca:
	//	return &references.BREAKPOINT{}
	//case 0xfe:
	//	return &extended.IMPLLIMENTATION_SPECIFIC{}
	//case 0xff:
	//	return &extended.EXTENDED{}
	default:
		panic(fmt.Errorf("Unsupported opcode: 0x%x!", opcode))
	}
	return nil
}
