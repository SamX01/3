package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var madd2_code cu.Function

type madd2_args struct {
	arg_dst  unsafe.Pointer
	arg_src1 unsafe.Pointer
	arg_fac1 float32
	arg_src2 unsafe.Pointer
	arg_fac2 float32
	arg_N    int
	argptr   [6]unsafe.Pointer
}

// Wrapper for madd2 CUDA kernel, asynchronous.
func k_madd2_async(dst unsafe.Pointer, src1 unsafe.Pointer, fac1 float32, src2 unsafe.Pointer, fac2 float32, N int, cfg *Config, str cu.Stream) {
	if madd2_code == 0 {
		madd2_code = cu.ModuleLoadData(madd2_ptx).GetFunction("madd2")
	}

	var a madd2_args

	a.arg_dst = dst
	a.argptr[0] = unsafe.Pointer(&a.arg_dst)
	a.arg_src1 = src1
	a.argptr[1] = unsafe.Pointer(&a.arg_src1)
	a.arg_fac1 = fac1
	a.argptr[2] = unsafe.Pointer(&a.arg_fac1)
	a.arg_src2 = src2
	a.argptr[3] = unsafe.Pointer(&a.arg_src2)
	a.arg_fac2 = fac2
	a.argptr[4] = unsafe.Pointer(&a.arg_fac2)
	a.arg_N = N
	a.argptr[5] = unsafe.Pointer(&a.arg_N)

	args := a.argptr[:]
	cu.LaunchKernel(madd2_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, str, args)
}

// Wrapper for madd2 CUDA kernel, synchronized.
func k_madd2(dst unsafe.Pointer, src1 unsafe.Pointer, fac1 float32, src2 unsafe.Pointer, fac2 float32, N int, cfg *Config) {
	str := Stream()
	k_madd2_async(dst, src1, fac1, src2, fac2, N, cfg, str)
	SyncAndRecycle(str)
}

const madd2_ptx = `
.version 3.0
.target sm_30
.address_size 64


.entry madd2(
	.param .u64 madd2_param_0,
	.param .u64 madd2_param_1,
	.param .f32 madd2_param_2,
	.param .u64 madd2_param_3,
	.param .f32 madd2_param_4,
	.param .u32 madd2_param_5
)
{
	.reg .f32 	%f<15>;
	.reg .pred 	%p<4>;
	.reg .s32 	%r<12>;
	.reg .s64 	%rl<15>;


	ld.param.u64 	%rl6, [madd2_param_0];
	ld.param.u64 	%rl1, [madd2_param_1];
	ld.param.u64 	%rl2, [madd2_param_3];
	ld.param.u32 	%r2, [madd2_param_5];
	cvta.to.global.u64 	%rl3, %rl6;
	cvta.to.global.u64 	%rl4, %rl2;
	cvta.to.global.u64 	%rl5, %rl1;
	.loc 2 9 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 2 11 1
	setp.lt.s32 	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	.loc 2 16 2
	ret;

BB0_2:
	ld.param.u64 	%rl13, [madd2_param_1];
	.loc 2 12 1
	setp.eq.s64 	%p2, %rl13, 0;
	@%p2 bra 	BB0_4;

	mul.wide.s32 	%rl7, %r1, 4;
	add.s64 	%rl8, %rl5, %rl7;
	ld.global.f32 	%f13, [%rl8];
	bra.uni 	BB0_5;

BB0_4:
	mov.f32 	%f13, 0f3F800000;

BB0_5:
	ld.param.u64 	%rl14, [madd2_param_3];
	.loc 2 13 1
	setp.eq.s64 	%p3, %rl14, 0;
	@%p3 bra 	BB0_7;

	mul.wide.s32 	%rl9, %r1, 4;
	add.s64 	%rl10, %rl4, %rl9;
	ld.global.f32 	%f14, [%rl10];
	bra.uni 	BB0_8;

BB0_7:
	mov.f32 	%f14, 0f3F800000;

BB0_8:
	ld.param.f32 	%f12, [madd2_param_4];
	.loc 2 14 1
	mul.f32 	%f9, %f14, %f12;
	ld.param.f32 	%f11, [madd2_param_2];
	.loc 2 14 1
	fma.rn.f32 	%f10, %f13, %f11, %f9;
	mul.wide.s32 	%rl11, %r1, 4;
	add.s64 	%rl12, %rl3, %rl11;
	st.global.f32 	[%rl12], %f10;
	.loc 2 16 2
	ret;
}


`
