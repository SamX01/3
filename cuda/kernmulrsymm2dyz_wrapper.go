package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var kernmulRSymm2Dyz_code cu.Function

type kernmulRSymm2Dyz_args struct {
	arg_fftMy  unsafe.Pointer
	arg_fftMz  unsafe.Pointer
	arg_fftKyy unsafe.Pointer
	arg_fftKzz unsafe.Pointer
	arg_fftKyz unsafe.Pointer
	arg_N1     int
	arg_N2     int
	argptr     [7]unsafe.Pointer
}

// Wrapper for kernmulRSymm2Dyz CUDA kernel, asynchronous.
func k_kernmulRSymm2Dyz_async(fftMy unsafe.Pointer, fftMz unsafe.Pointer, fftKyy unsafe.Pointer, fftKzz unsafe.Pointer, fftKyz unsafe.Pointer, N1 int, N2 int, cfg *Config, str cu.Stream) {
	if kernmulRSymm2Dyz_code == 0 {
		kernmulRSymm2Dyz_code = cu.ModuleLoadData(kernmulRSymm2Dyz_ptx).GetFunction("kernmulRSymm2Dyz")
	}

	var a kernmulRSymm2Dyz_args

	a.arg_fftMy = fftMy
	a.argptr[0] = unsafe.Pointer(&a.arg_fftMy)
	a.arg_fftMz = fftMz
	a.argptr[1] = unsafe.Pointer(&a.arg_fftMz)
	a.arg_fftKyy = fftKyy
	a.argptr[2] = unsafe.Pointer(&a.arg_fftKyy)
	a.arg_fftKzz = fftKzz
	a.argptr[3] = unsafe.Pointer(&a.arg_fftKzz)
	a.arg_fftKyz = fftKyz
	a.argptr[4] = unsafe.Pointer(&a.arg_fftKyz)
	a.arg_N1 = N1
	a.argptr[5] = unsafe.Pointer(&a.arg_N1)
	a.arg_N2 = N2
	a.argptr[6] = unsafe.Pointer(&a.arg_N2)

	args := a.argptr[:]
	cu.LaunchKernel(kernmulRSymm2Dyz_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, str, args)
}

// Wrapper for kernmulRSymm2Dyz CUDA kernel, synchronized.
func k_kernmulRSymm2Dyz(fftMy unsafe.Pointer, fftMz unsafe.Pointer, fftKyy unsafe.Pointer, fftKzz unsafe.Pointer, fftKyz unsafe.Pointer, N1 int, N2 int, cfg *Config) {
	str := Stream()
	k_kernmulRSymm2Dyz_async(fftMy, fftMz, fftKyy, fftKzz, fftKyz, N1, N2, cfg, str)
	SyncAndRecycle(str)
}

const kernmulRSymm2Dyz_ptx = `
.version 3.0
.target sm_30
.address_size 64


.entry kernmulRSymm2Dyz(
	.param .u64 kernmulRSymm2Dyz_param_0,
	.param .u64 kernmulRSymm2Dyz_param_1,
	.param .u64 kernmulRSymm2Dyz_param_2,
	.param .u64 kernmulRSymm2Dyz_param_3,
	.param .u64 kernmulRSymm2Dyz_param_4,
	.param .u32 kernmulRSymm2Dyz_param_5,
	.param .u32 kernmulRSymm2Dyz_param_6
)
{
	.reg .f32 	%f<20>;
	.reg .pred 	%p<5>;
	.reg .s32 	%r<36>;
	.reg .s64 	%rl<28>;


	ld.param.u64 	%rl10, [kernmulRSymm2Dyz_param_0];
	ld.param.u64 	%rl11, [kernmulRSymm2Dyz_param_1];
	ld.param.u64 	%rl12, [kernmulRSymm2Dyz_param_4];
	ld.param.u32 	%r1, [kernmulRSymm2Dyz_param_5];
	ld.param.u32 	%r2, [kernmulRSymm2Dyz_param_6];
	cvta.to.global.u64 	%rl3, %rl12;
	cvta.to.global.u64 	%rl4, %rl11;
	cvta.to.global.u64 	%rl5, %rl10;
	.loc 2 29 1
	mov.u32 	%r8, %ntid.y;
	mov.u32 	%r9, %ctaid.y;
	mov.u32 	%r10, %tid.y;
	mad.lo.s32 	%r3, %r8, %r9, %r10;
	.loc 2 30 1
	mov.u32 	%r11, %ntid.x;
	mov.u32 	%r12, %ctaid.x;
	mov.u32 	%r13, %tid.x;
	mad.lo.s32 	%r4, %r11, %r12, %r13;
	.loc 2 32 1
	setp.ge.s32 	%p1, %r4, %r2;
	setp.ge.s32 	%p2, %r3, %r1;
	or.pred  	%p3, %p1, %p2;
	.loc 2 32 1
	@%p3 bra 	BB0_5;

	ld.param.u32 	%r34, [kernmulRSymm2Dyz_param_6];
	.loc 2 36 1
	mad.lo.s32 	%r5, %r3, %r34, %r4;
	ld.param.u32 	%r33, [kernmulRSymm2Dyz_param_5];
	.loc 2 37 1
	sub.s32 	%r14, %r33, %r3;
	mad.lo.s32 	%r6, %r14, %r34, %r4;
	shl.b32 	%r15, %r5, 1;
	.loc 2 41 1
	mul.wide.s32 	%rl13, %r15, 4;
	add.s64 	%rl6, %rl5, %rl13;
	ld.global.f32 	%f1, [%rl6];
	.loc 2 42 1
	or.b32  	%r17, %r15, 1;
	mul.wide.s32 	%rl14, %r17, 4;
	add.s64 	%rl15, %rl5, %rl14;
	add.s64 	%rl7, %rl15, -4;
	.loc 2 42 1
	ld.global.f32 	%f2, [%rl15];
	.loc 2 43 1
	add.s64 	%rl8, %rl4, %rl13;
	ld.global.f32 	%f3, [%rl8];
	.loc 2 44 1
	add.s64 	%rl16, %rl4, %rl14;
	add.s64 	%rl9, %rl16, -4;
	.loc 2 44 1
	ld.global.f32 	%f4, [%rl16];
	.loc 2 47 1
	shr.u32 	%r21, %r33, 31;
	add.s32 	%r22, %r33, %r21;
	shr.s32 	%r23, %r22, 1;
	add.s32 	%r24, %r23, 1;
	setp.lt.s32 	%p4, %r3, %r24;
	@%p4 bra 	BB0_3;

	.loc 2 54 1
	mul.wide.s32 	%rl17, %r6, 4;
	add.s64 	%rl18, %rl3, %rl17;
	ld.global.f32 	%f8, [%rl18];
	neg.f32 	%f19, %f8;
	mov.u32 	%r35, %r6;
	bra.uni 	BB0_4;

BB0_3:
	.loc 2 50 1
	mul.wide.s32 	%rl19, %r5, 4;
	add.s64 	%rl20, %rl3, %rl19;
	ld.global.f32 	%f19, [%rl20];
	mov.u32 	%r35, %r5;

BB0_4:
	ld.param.u64 	%rl27, [kernmulRSymm2Dyz_param_3];
	cvta.to.global.u64 	%rl21, %rl27;
	mul.wide.s32 	%rl22, %r35, 4;
	add.s64 	%rl23, %rl21, %rl22;
	ld.param.u64 	%rl26, [kernmulRSymm2Dyz_param_2];
	cvta.to.global.u64 	%rl24, %rl26;
	add.s64 	%rl25, %rl24, %rl22;
	ld.global.f32 	%f9, [%rl23];
	ld.global.f32 	%f10, [%rl25];
	.loc 2 57 1
	mul.f32 	%f11, %f3, %f19;
	fma.rn.f32 	%f12, %f1, %f10, %f11;
	st.global.f32 	[%rl6], %f12;
	.loc 2 58 1
	mul.f32 	%f13, %f4, %f19;
	fma.rn.f32 	%f14, %f2, %f10, %f13;
	st.global.f32 	[%rl7+4], %f14;
	.loc 2 59 1
	mul.f32 	%f15, %f3, %f9;
	fma.rn.f32 	%f16, %f1, %f19, %f15;
	st.global.f32 	[%rl8], %f16;
	.loc 2 60 1
	mul.f32 	%f17, %f4, %f9;
	fma.rn.f32 	%f18, %f2, %f19, %f17;
	st.global.f32 	[%rl9+4], %f18;

BB0_5:
	.loc 2 61 2
	ret;
}


`
