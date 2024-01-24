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


TEXT ·SimdPopcntQuad(SB),NOSPLIT,$0-64
    VMOVDQU nums+0(FP), Y0 // Y0 = x
    MOVQ $0x5555555555555555, AX
    MOVQ AX, X9
    VPBROADCASTQ X9, Y5 // Y5 = m0
    MOVQ $0x3333333333333333, AX
    MOVQ AX, X9
    VPBROADCASTQ X9, Y6 // Y6 = m1
    MOVQ $0x0f0f0f0f0f0f0f0f, AX
    MOVQ AX, X9
    VPBROADCASTQ X9, Y7 // Y7 = m2
    MOVQ $0x7f, AX
    MOVQ AX, X9
    VPBROADCASTQ X9, Y8 // Y8 = m
    VPSRLQ $1, Y0, Y1 // Y1 = x>>1
    VPAND Y1, Y5, Y1 // Y1 = x>>1&m0
    VPAND Y0, Y5, Y2 // Y2 = x&m0
    VPADDQ Y1, Y2, Y0 // x = x>>1&m0 + x&m0
    VPSRLQ $2, Y0, Y1 // Y1 = x>>2
    VPAND Y1, Y6, Y1 // Y1 = x>>2&m1
    VPAND Y0, Y6, Y2 // Y2 = x&m1
    VPADDQ Y1, Y2, Y0 // x = x>>2&m1 + x&m1
    VPSRLQ $4, Y0, Y1 // Y1 = x>>4
    VPAND Y1, Y7, Y1 // Y1 = x>>4&m2
    VPAND Y0, Y7, Y2 // Y2 = x&m2
    VPADDQ Y1, Y2, Y0 // x = x>>2&m2 + x&m2
    VPSRLQ $8, Y0, Y1 // Y1 = x >> 8
    VPADDQ Y1, Y0, Y0 // x += x >> 8
    VPSRLQ $16, Y0, Y1 // Y1 = x >> 16
    VPADDQ Y1, Y0, Y0 // x += x >> 16
    VPSRLQ $32, Y0, Y1 // Y1 = x >> 32
    VPADDQ Y1, Y0, Y0 // x += x >> 32
    VPAND Y0, Y8, Y0 // x & (1<<7-1)
    VMOVDQU Y0, ret+32(FP)
    RET
