#include "textflag.h"

TEXT ·shiftRight1Quad(SB),NOSPLIT,$0-64
    VMOVDQU nums+0(FP), Y0
    VPSRLQ $1, Y0, Y0
    VMOVDQU Y0, ret+32(FP)
    RET

TEXT ·shiftRight2Quad(SB),NOSPLIT,$0-64
    VMOVDQU nums+0(FP), Y0
    VPSRLQ $2, Y0, Y0
    VMOVDQU Y0, ret+32(FP)
    RET

TEXT ·shiftRight4Quad(SB),NOSPLIT,$0-64
    VMOVDQU nums+0(FP), Y0
    VPSRLQ $4, Y0, Y0
    VMOVDQU Y0, ret+32(FP)
    RET

TEXT ·shiftRight8Quad(SB),NOSPLIT,$0-64
    VMOVDQU nums+0(FP), Y0
    VPSRLQ $8, Y0, Y0
    VMOVDQU Y0, ret+32(FP)
    RET

TEXT ·shiftRight16Quad(SB),NOSPLIT,$0-64
    VMOVDQU nums+0(FP), Y0
    VPSRLQ $16, Y0, Y0
    VMOVDQU Y0, ret+32(FP)
    RET

TEXT ·shiftRight32Quad(SB),NOSPLIT,$0-64
    VMOVDQU nums+0(FP), Y0
    VPSRLQ $32, Y0, Y0
    VMOVDQU Y0, ret+32(FP)
    RET

TEXT ·andQuad(SB),NOSPLIT,$0-96
    VMOVDQU a+0(FP), Y0
    VMOVDQU b+32(FP), Y1
    VPAND Y0, Y1, Y2
    VMOVDQU Y2, ret+64(FP)
    RET

TEXT ·addQuad(SB),NOSPLIT,$0-96
    VMOVDQU a+0(FP), Y0
    VMOVDQU b+32(FP), Y1
    VPADDQ Y0, Y1, Y2
    VMOVDQU Y2, ret+64(FP)
    RET
