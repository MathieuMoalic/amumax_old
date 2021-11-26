package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/MathieuMoalic/amumax/cuda/cu"
	"github.com/MathieuMoalic/amumax/timer"
	"sync"
	"unsafe"
)

// CUDA handle for kernmulRSymm2Dxy kernel
var kernmulRSymm2Dxy_code cu.Function

// Stores the arguments for kernmulRSymm2Dxy kernel invocation
type kernmulRSymm2Dxy_args_t struct {
	arg_fftMx  unsafe.Pointer
	arg_fftMy  unsafe.Pointer
	arg_fftKxx unsafe.Pointer
	arg_fftKyy unsafe.Pointer
	arg_fftKxy unsafe.Pointer
	arg_Nx     int
	arg_Ny     int
	argptr     [7]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for kernmulRSymm2Dxy kernel invocation
var kernmulRSymm2Dxy_args kernmulRSymm2Dxy_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	kernmulRSymm2Dxy_args.argptr[0] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_fftMx)
	kernmulRSymm2Dxy_args.argptr[1] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_fftMy)
	kernmulRSymm2Dxy_args.argptr[2] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_fftKxx)
	kernmulRSymm2Dxy_args.argptr[3] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_fftKyy)
	kernmulRSymm2Dxy_args.argptr[4] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_fftKxy)
	kernmulRSymm2Dxy_args.argptr[5] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_Nx)
	kernmulRSymm2Dxy_args.argptr[6] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_Ny)
}

// Wrapper for kernmulRSymm2Dxy CUDA kernel, asynchronous.
func k_kernmulRSymm2Dxy_async(fftMx unsafe.Pointer, fftMy unsafe.Pointer, fftKxx unsafe.Pointer, fftKyy unsafe.Pointer, fftKxy unsafe.Pointer, Nx int, Ny int, cfg *config) {
	if Synchronous { // debug
		Sync()
		timer.Start("kernmulRSymm2Dxy")
	}

	kernmulRSymm2Dxy_args.Lock()
	defer kernmulRSymm2Dxy_args.Unlock()

	if kernmulRSymm2Dxy_code == 0 {
		kernmulRSymm2Dxy_code = fatbinLoad(kernmulRSymm2Dxy_map, "kernmulRSymm2Dxy")
	}

	kernmulRSymm2Dxy_args.arg_fftMx = fftMx
	kernmulRSymm2Dxy_args.arg_fftMy = fftMy
	kernmulRSymm2Dxy_args.arg_fftKxx = fftKxx
	kernmulRSymm2Dxy_args.arg_fftKyy = fftKyy
	kernmulRSymm2Dxy_args.arg_fftKxy = fftKxy
	kernmulRSymm2Dxy_args.arg_Nx = Nx
	kernmulRSymm2Dxy_args.arg_Ny = Ny

	args := kernmulRSymm2Dxy_args.argptr[:]
	cu.LaunchKernel(kernmulRSymm2Dxy_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
		timer.Stop("kernmulRSymm2Dxy")
	}
}

// maps compute capability on PTX code for kernmulRSymm2Dxy kernel.
var kernmulRSymm2Dxy_map = map[int]string{0: "",
	30: kernmulRSymm2Dxy_ptx_30,
	32: kernmulRSymm2Dxy_ptx_32,
	35: kernmulRSymm2Dxy_ptx_35,
	37: kernmulRSymm2Dxy_ptx_37,
	50: kernmulRSymm2Dxy_ptx_50,
	52: kernmulRSymm2Dxy_ptx_52,
	53: kernmulRSymm2Dxy_ptx_53,
	60: kernmulRSymm2Dxy_ptx_60,
	61: kernmulRSymm2Dxy_ptx_61,
	62: kernmulRSymm2Dxy_ptx_62,
	70: kernmulRSymm2Dxy_ptx_70,
	72: kernmulRSymm2Dxy_ptx_72,
	75: kernmulRSymm2Dxy_ptx_75}

// kernmulRSymm2Dxy PTX code for various compute capabilities.
const (
	kernmulRSymm2Dxy_ptx_30 = `
.version 6.5
.target sm_30
.address_size 64

	// .globl	kernmulRSymm2Dxy

.visible .entry kernmulRSymm2Dxy(
	.param .u64 kernmulRSymm2Dxy_param_0,
	.param .u64 kernmulRSymm2Dxy_param_1,
	.param .u64 kernmulRSymm2Dxy_param_2,
	.param .u64 kernmulRSymm2Dxy_param_3,
	.param .u64 kernmulRSymm2Dxy_param_4,
	.param .u32 kernmulRSymm2Dxy_param_5,
	.param .u32 kernmulRSymm2Dxy_param_6
)
{
	.reg .pred 	%p<5>;
	.reg .f32 	%f<18>;
	.reg .b32 	%r<19>;
	.reg .b64 	%rd<18>;


	ld.param.u64 	%rd1, [kernmulRSymm2Dxy_param_0];
	ld.param.u64 	%rd2, [kernmulRSymm2Dxy_param_1];
	ld.param.u64 	%rd3, [kernmulRSymm2Dxy_param_2];
	ld.param.u64 	%rd4, [kernmulRSymm2Dxy_param_3];
	ld.param.u64 	%rd5, [kernmulRSymm2Dxy_param_4];
	ld.param.u32 	%r3, [kernmulRSymm2Dxy_param_5];
	ld.param.u32 	%r4, [kernmulRSymm2Dxy_param_6];
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r1, %r5, %r6, %r7;
	mov.u32 	%r8, %ntid.y;
	mov.u32 	%r9, %ctaid.y;
	mov.u32 	%r10, %tid.y;
	mad.lo.s32 	%r2, %r8, %r9, %r10;
	setp.ge.s32	%p1, %r2, %r4;
	setp.ge.s32	%p2, %r1, %r3;
	or.pred  	%p3, %p1, %p2;
	@%p3 bra 	BB0_2;

	cvta.to.global.u64 	%rd6, %rd2;
	cvta.to.global.u64 	%rd7, %rd1;
	cvta.to.global.u64 	%rd8, %rd3;
	mad.lo.s32 	%r11, %r2, %r3, %r1;
	shl.b32 	%r12, %r11, 1;
	mul.wide.s32 	%rd9, %r12, 4;
	add.s64 	%rd10, %rd7, %rd9;
	ld.global.f32 	%f1, [%rd10+4];
	add.s64 	%rd11, %rd6, %rd9;
	ld.global.f32 	%f2, [%rd11+4];
	shr.u32 	%r13, %r4, 31;
	add.s32 	%r14, %r4, %r13;
	shr.s32 	%r15, %r14, 1;
	setp.gt.s32	%p4, %r2, %r15;
	sub.s32 	%r16, %r4, %r2;
	selp.b32	%r17, %r16, %r2, %p4;
	selp.f32	%f3, 0fBF800000, 0f3F800000, %p4;
	mad.lo.s32 	%r18, %r17, %r3, %r1;
	mul.wide.s32 	%rd12, %r18, 4;
	add.s64 	%rd13, %rd8, %rd12;
	cvta.to.global.u64 	%rd14, %rd4;
	add.s64 	%rd15, %rd14, %rd12;
	ld.global.f32 	%f4, [%rd15];
	cvta.to.global.u64 	%rd16, %rd5;
	add.s64 	%rd17, %rd16, %rd12;
	ld.global.f32 	%f5, [%rd17];
	mul.f32 	%f6, %f3, %f5;
	ld.global.f32 	%f7, [%rd13];
	ld.global.f32 	%f8, [%rd10];
	ld.global.f32 	%f9, [%rd11];
	mul.f32 	%f10, %f9, %f6;
	fma.rn.f32 	%f11, %f8, %f7, %f10;
	st.global.f32 	[%rd10], %f11;
	mul.f32 	%f12, %f2, %f6;
	fma.rn.f32 	%f13, %f1, %f7, %f12;
	st.global.f32 	[%rd10+4], %f13;
	mul.f32 	%f14, %f8, %f6;
	fma.rn.f32 	%f15, %f9, %f4, %f14;
	st.global.f32 	[%rd11], %f15;
	mul.f32 	%f16, %f1, %f6;
	fma.rn.f32 	%f17, %f2, %f4, %f16;
	st.global.f32 	[%rd11+4], %f17;

BB0_2:
	ret;
}


`
	kernmulRSymm2Dxy_ptx_32 = `
.version 6.5
.target sm_32
.address_size 64

	// .globl	kernmulRSymm2Dxy

.visible .entry kernmulRSymm2Dxy(
	.param .u64 kernmulRSymm2Dxy_param_0,
	.param .u64 kernmulRSymm2Dxy_param_1,
	.param .u64 kernmulRSymm2Dxy_param_2,
	.param .u64 kernmulRSymm2Dxy_param_3,
	.param .u64 kernmulRSymm2Dxy_param_4,
	.param .u32 kernmulRSymm2Dxy_param_5,
	.param .u32 kernmulRSymm2Dxy_param_6
)
{
	.reg .pred 	%p<5>;
	.reg .f32 	%f<18>;
	.reg .b32 	%r<19>;
	.reg .b64 	%rd<18>;


	ld.param.u64 	%rd1, [kernmulRSymm2Dxy_param_0];
	ld.param.u64 	%rd2, [kernmulRSymm2Dxy_param_1];
	ld.param.u64 	%rd3, [kernmulRSymm2Dxy_param_2];
	ld.param.u64 	%rd4, [kernmulRSymm2Dxy_param_3];
	ld.param.u64 	%rd5, [kernmulRSymm2Dxy_param_4];
	ld.param.u32 	%r3, [kernmulRSymm2Dxy_param_5];
	ld.param.u32 	%r4, [kernmulRSymm2Dxy_param_6];
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r1, %r5, %r6, %r7;
	mov.u32 	%r8, %ntid.y;
	mov.u32 	%r9, %ctaid.y;
	mov.u32 	%r10, %tid.y;
	mad.lo.s32 	%r2, %r8, %r9, %r10;
	setp.ge.s32	%p1, %r2, %r4;
	setp.ge.s32	%p2, %r1, %r3;
	or.pred  	%p3, %p1, %p2;
	@%p3 bra 	BB0_2;

	cvta.to.global.u64 	%rd6, %rd2;
	cvta.to.global.u64 	%rd7, %rd1;
	cvta.to.global.u64 	%rd8, %rd3;
	mad.lo.s32 	%r11, %r2, %r3, %r1;
	shl.b32 	%r12, %r11, 1;
	mul.wide.s32 	%rd9, %r12, 4;
	add.s64 	%rd10, %rd7, %rd9;
	ld.global.f32 	%f1, [%rd10+4];
	add.s64 	%rd11, %rd6, %rd9;
	ld.global.f32 	%f2, [%rd11+4];
	shr.u32 	%r13, %r4, 31;
	add.s32 	%r14, %r4, %r13;
	shr.s32 	%r15, %r14, 1;
	setp.gt.s32	%p4, %r2, %r15;
	sub.s32 	%r16, %r4, %r2;
	selp.b32	%r17, %r16, %r2, %p4;
	selp.f32	%f3, 0fBF800000, 0f3F800000, %p4;
	mad.lo.s32 	%r18, %r17, %r3, %r1;
	mul.wide.s32 	%rd12, %r18, 4;
	add.s64 	%rd13, %rd8, %rd12;
	cvta.to.global.u64 	%rd14, %rd4;
	add.s64 	%rd15, %rd14, %rd12;
	ld.global.nc.f32 	%f4, [%rd15];
	cvta.to.global.u64 	%rd16, %rd5;
	add.s64 	%rd17, %rd16, %rd12;
	ld.global.nc.f32 	%f5, [%rd17];
	mul.f32 	%f6, %f3, %f5;
	ld.global.nc.f32 	%f7, [%rd13];
	ld.global.f32 	%f8, [%rd10];
	ld.global.f32 	%f9, [%rd11];
	mul.f32 	%f10, %f9, %f6;
	fma.rn.f32 	%f11, %f8, %f7, %f10;
	st.global.f32 	[%rd10], %f11;
	mul.f32 	%f12, %f2, %f6;
	fma.rn.f32 	%f13, %f1, %f7, %f12;
	st.global.f32 	[%rd10+4], %f13;
	mul.f32 	%f14, %f8, %f6;
	fma.rn.f32 	%f15, %f9, %f4, %f14;
	st.global.f32 	[%rd11], %f15;
	mul.f32 	%f16, %f1, %f6;
	fma.rn.f32 	%f17, %f2, %f4, %f16;
	st.global.f32 	[%rd11+4], %f17;

BB0_2:
	ret;
}


`
	kernmulRSymm2Dxy_ptx_35 = `
.version 6.5
.target sm_35
.address_size 64

	// .globl	kernmulRSymm2Dxy

.visible .entry kernmulRSymm2Dxy(
	.param .u64 kernmulRSymm2Dxy_param_0,
	.param .u64 kernmulRSymm2Dxy_param_1,
	.param .u64 kernmulRSymm2Dxy_param_2,
	.param .u64 kernmulRSymm2Dxy_param_3,
	.param .u64 kernmulRSymm2Dxy_param_4,
	.param .u32 kernmulRSymm2Dxy_param_5,
	.param .u32 kernmulRSymm2Dxy_param_6
)
{
	.reg .pred 	%p<5>;
	.reg .f32 	%f<18>;
	.reg .b32 	%r<19>;
	.reg .b64 	%rd<18>;


	ld.param.u64 	%rd1, [kernmulRSymm2Dxy_param_0];
	ld.param.u64 	%rd2, [kernmulRSymm2Dxy_param_1];
	ld.param.u64 	%rd3, [kernmulRSymm2Dxy_param_2];
	ld.param.u64 	%rd4, [kernmulRSymm2Dxy_param_3];
	ld.param.u64 	%rd5, [kernmulRSymm2Dxy_param_4];
	ld.param.u32 	%r3, [kernmulRSymm2Dxy_param_5];
	ld.param.u32 	%r4, [kernmulRSymm2Dxy_param_6];
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r1, %r5, %r6, %r7;
	mov.u32 	%r8, %ntid.y;
	mov.u32 	%r9, %ctaid.y;
	mov.u32 	%r10, %tid.y;
	mad.lo.s32 	%r2, %r8, %r9, %r10;
	setp.ge.s32	%p1, %r2, %r4;
	setp.ge.s32	%p2, %r1, %r3;
	or.pred  	%p3, %p1, %p2;
	@%p3 bra 	BB0_2;

	cvta.to.global.u64 	%rd6, %rd2;
	cvta.to.global.u64 	%rd7, %rd1;
	cvta.to.global.u64 	%rd8, %rd3;
	mad.lo.s32 	%r11, %r2, %r3, %r1;
	shl.b32 	%r12, %r11, 1;
	mul.wide.s32 	%rd9, %r12, 4;
	add.s64 	%rd10, %rd7, %rd9;
	ld.global.f32 	%f1, [%rd10+4];
	add.s64 	%rd11, %rd6, %rd9;
	ld.global.f32 	%f2, [%rd11+4];
	shr.u32 	%r13, %r4, 31;
	add.s32 	%r14, %r4, %r13;
	shr.s32 	%r15, %r14, 1;
	setp.gt.s32	%p4, %r2, %r15;
	sub.s32 	%r16, %r4, %r2;
	selp.b32	%r17, %r16, %r2, %p4;
	selp.f32	%f3, 0fBF800000, 0f3F800000, %p4;
	mad.lo.s32 	%r18, %r17, %r3, %r1;
	mul.wide.s32 	%rd12, %r18, 4;
	add.s64 	%rd13, %rd8, %rd12;
	cvta.to.global.u64 	%rd14, %rd4;
	add.s64 	%rd15, %rd14, %rd12;
	ld.global.nc.f32 	%f4, [%rd15];
	cvta.to.global.u64 	%rd16, %rd5;
	add.s64 	%rd17, %rd16, %rd12;
	ld.global.nc.f32 	%f5, [%rd17];
	mul.f32 	%f6, %f3, %f5;
	ld.global.nc.f32 	%f7, [%rd13];
	ld.global.f32 	%f8, [%rd10];
	ld.global.f32 	%f9, [%rd11];
	mul.f32 	%f10, %f9, %f6;
	fma.rn.f32 	%f11, %f8, %f7, %f10;
	st.global.f32 	[%rd10], %f11;
	mul.f32 	%f12, %f2, %f6;
	fma.rn.f32 	%f13, %f1, %f7, %f12;
	st.global.f32 	[%rd10+4], %f13;
	mul.f32 	%f14, %f8, %f6;
	fma.rn.f32 	%f15, %f9, %f4, %f14;
	st.global.f32 	[%rd11], %f15;
	mul.f32 	%f16, %f1, %f6;
	fma.rn.f32 	%f17, %f2, %f4, %f16;
	st.global.f32 	[%rd11+4], %f17;

BB0_2:
	ret;
}


`
	kernmulRSymm2Dxy_ptx_37 = `
.version 6.5
.target sm_37
.address_size 64

	// .globl	kernmulRSymm2Dxy

.visible .entry kernmulRSymm2Dxy(
	.param .u64 kernmulRSymm2Dxy_param_0,
	.param .u64 kernmulRSymm2Dxy_param_1,
	.param .u64 kernmulRSymm2Dxy_param_2,
	.param .u64 kernmulRSymm2Dxy_param_3,
	.param .u64 kernmulRSymm2Dxy_param_4,
	.param .u32 kernmulRSymm2Dxy_param_5,
	.param .u32 kernmulRSymm2Dxy_param_6
)
{
	.reg .pred 	%p<5>;
	.reg .f32 	%f<18>;
	.reg .b32 	%r<19>;
	.reg .b64 	%rd<18>;


	ld.param.u64 	%rd1, [kernmulRSymm2Dxy_param_0];
	ld.param.u64 	%rd2, [kernmulRSymm2Dxy_param_1];
	ld.param.u64 	%rd3, [kernmulRSymm2Dxy_param_2];
	ld.param.u64 	%rd4, [kernmulRSymm2Dxy_param_3];
	ld.param.u64 	%rd5, [kernmulRSymm2Dxy_param_4];
	ld.param.u32 	%r3, [kernmulRSymm2Dxy_param_5];
	ld.param.u32 	%r4, [kernmulRSymm2Dxy_param_6];
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r1, %r5, %r6, %r7;
	mov.u32 	%r8, %ntid.y;
	mov.u32 	%r9, %ctaid.y;
	mov.u32 	%r10, %tid.y;
	mad.lo.s32 	%r2, %r8, %r9, %r10;
	setp.ge.s32	%p1, %r2, %r4;
	setp.ge.s32	%p2, %r1, %r3;
	or.pred  	%p3, %p1, %p2;
	@%p3 bra 	BB0_2;

	cvta.to.global.u64 	%rd6, %rd2;
	cvta.to.global.u64 	%rd7, %rd1;
	cvta.to.global.u64 	%rd8, %rd3;
	mad.lo.s32 	%r11, %r2, %r3, %r1;
	shl.b32 	%r12, %r11, 1;
	mul.wide.s32 	%rd9, %r12, 4;
	add.s64 	%rd10, %rd7, %rd9;
	ld.global.f32 	%f1, [%rd10+4];
	add.s64 	%rd11, %rd6, %rd9;
	ld.global.f32 	%f2, [%rd11+4];
	shr.u32 	%r13, %r4, 31;
	add.s32 	%r14, %r4, %r13;
	shr.s32 	%r15, %r14, 1;
	setp.gt.s32	%p4, %r2, %r15;
	sub.s32 	%r16, %r4, %r2;
	selp.b32	%r17, %r16, %r2, %p4;
	selp.f32	%f3, 0fBF800000, 0f3F800000, %p4;
	mad.lo.s32 	%r18, %r17, %r3, %r1;
	mul.wide.s32 	%rd12, %r18, 4;
	add.s64 	%rd13, %rd8, %rd12;
	cvta.to.global.u64 	%rd14, %rd4;
	add.s64 	%rd15, %rd14, %rd12;
	ld.global.nc.f32 	%f4, [%rd15];
	cvta.to.global.u64 	%rd16, %rd5;
	add.s64 	%rd17, %rd16, %rd12;
	ld.global.nc.f32 	%f5, [%rd17];
	mul.f32 	%f6, %f3, %f5;
	ld.global.nc.f32 	%f7, [%rd13];
	ld.global.f32 	%f8, [%rd10];
	ld.global.f32 	%f9, [%rd11];
	mul.f32 	%f10, %f9, %f6;
	fma.rn.f32 	%f11, %f8, %f7, %f10;
	st.global.f32 	[%rd10], %f11;
	mul.f32 	%f12, %f2, %f6;
	fma.rn.f32 	%f13, %f1, %f7, %f12;
	st.global.f32 	[%rd10+4], %f13;
	mul.f32 	%f14, %f8, %f6;
	fma.rn.f32 	%f15, %f9, %f4, %f14;
	st.global.f32 	[%rd11], %f15;
	mul.f32 	%f16, %f1, %f6;
	fma.rn.f32 	%f17, %f2, %f4, %f16;
	st.global.f32 	[%rd11+4], %f17;

BB0_2:
	ret;
}


`
	kernmulRSymm2Dxy_ptx_50 = `
.version 6.5
.target sm_50
.address_size 64

	// .globl	kernmulRSymm2Dxy

.visible .entry kernmulRSymm2Dxy(
	.param .u64 kernmulRSymm2Dxy_param_0,
	.param .u64 kernmulRSymm2Dxy_param_1,
	.param .u64 kernmulRSymm2Dxy_param_2,
	.param .u64 kernmulRSymm2Dxy_param_3,
	.param .u64 kernmulRSymm2Dxy_param_4,
	.param .u32 kernmulRSymm2Dxy_param_5,
	.param .u32 kernmulRSymm2Dxy_param_6
)
{
	.reg .pred 	%p<5>;
	.reg .f32 	%f<18>;
	.reg .b32 	%r<19>;
	.reg .b64 	%rd<18>;


	ld.param.u64 	%rd1, [kernmulRSymm2Dxy_param_0];
	ld.param.u64 	%rd2, [kernmulRSymm2Dxy_param_1];
	ld.param.u64 	%rd3, [kernmulRSymm2Dxy_param_2];
	ld.param.u64 	%rd4, [kernmulRSymm2Dxy_param_3];
	ld.param.u64 	%rd5, [kernmulRSymm2Dxy_param_4];
	ld.param.u32 	%r3, [kernmulRSymm2Dxy_param_5];
	ld.param.u32 	%r4, [kernmulRSymm2Dxy_param_6];
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r1, %r5, %r6, %r7;
	mov.u32 	%r8, %ntid.y;
	mov.u32 	%r9, %ctaid.y;
	mov.u32 	%r10, %tid.y;
	mad.lo.s32 	%r2, %r8, %r9, %r10;
	setp.ge.s32	%p1, %r2, %r4;
	setp.ge.s32	%p2, %r1, %r3;
	or.pred  	%p3, %p1, %p2;
	@%p3 bra 	BB0_2;

	cvta.to.global.u64 	%rd6, %rd2;
	cvta.to.global.u64 	%rd7, %rd1;
	cvta.to.global.u64 	%rd8, %rd3;
	mad.lo.s32 	%r11, %r2, %r3, %r1;
	shl.b32 	%r12, %r11, 1;
	mul.wide.s32 	%rd9, %r12, 4;
	add.s64 	%rd10, %rd7, %rd9;
	ld.global.f32 	%f1, [%rd10+4];
	add.s64 	%rd11, %rd6, %rd9;
	ld.global.f32 	%f2, [%rd11+4];
	shr.u32 	%r13, %r4, 31;
	add.s32 	%r14, %r4, %r13;
	shr.s32 	%r15, %r14, 1;
	setp.gt.s32	%p4, %r2, %r15;
	sub.s32 	%r16, %r4, %r2;
	selp.b32	%r17, %r16, %r2, %p4;
	selp.f32	%f3, 0fBF800000, 0f3F800000, %p4;
	mad.lo.s32 	%r18, %r17, %r3, %r1;
	mul.wide.s32 	%rd12, %r18, 4;
	add.s64 	%rd13, %rd8, %rd12;
	cvta.to.global.u64 	%rd14, %rd4;
	add.s64 	%rd15, %rd14, %rd12;
	ld.global.nc.f32 	%f4, [%rd15];
	cvta.to.global.u64 	%rd16, %rd5;
	add.s64 	%rd17, %rd16, %rd12;
	ld.global.nc.f32 	%f5, [%rd17];
	mul.f32 	%f6, %f3, %f5;
	ld.global.nc.f32 	%f7, [%rd13];
	ld.global.f32 	%f8, [%rd10];
	ld.global.f32 	%f9, [%rd11];
	mul.f32 	%f10, %f9, %f6;
	fma.rn.f32 	%f11, %f8, %f7, %f10;
	st.global.f32 	[%rd10], %f11;
	mul.f32 	%f12, %f2, %f6;
	fma.rn.f32 	%f13, %f1, %f7, %f12;
	st.global.f32 	[%rd10+4], %f13;
	mul.f32 	%f14, %f8, %f6;
	fma.rn.f32 	%f15, %f9, %f4, %f14;
	st.global.f32 	[%rd11], %f15;
	mul.f32 	%f16, %f1, %f6;
	fma.rn.f32 	%f17, %f2, %f4, %f16;
	st.global.f32 	[%rd11+4], %f17;

BB0_2:
	ret;
}


`
	kernmulRSymm2Dxy_ptx_52 = `
.version 6.5
.target sm_52
.address_size 64

	// .globl	kernmulRSymm2Dxy

.visible .entry kernmulRSymm2Dxy(
	.param .u64 kernmulRSymm2Dxy_param_0,
	.param .u64 kernmulRSymm2Dxy_param_1,
	.param .u64 kernmulRSymm2Dxy_param_2,
	.param .u64 kernmulRSymm2Dxy_param_3,
	.param .u64 kernmulRSymm2Dxy_param_4,
	.param .u32 kernmulRSymm2Dxy_param_5,
	.param .u32 kernmulRSymm2Dxy_param_6
)
{
	.reg .pred 	%p<5>;
	.reg .f32 	%f<18>;
	.reg .b32 	%r<19>;
	.reg .b64 	%rd<18>;


	ld.param.u64 	%rd1, [kernmulRSymm2Dxy_param_0];
	ld.param.u64 	%rd2, [kernmulRSymm2Dxy_param_1];
	ld.param.u64 	%rd3, [kernmulRSymm2Dxy_param_2];
	ld.param.u64 	%rd4, [kernmulRSymm2Dxy_param_3];
	ld.param.u64 	%rd5, [kernmulRSymm2Dxy_param_4];
	ld.param.u32 	%r3, [kernmulRSymm2Dxy_param_5];
	ld.param.u32 	%r4, [kernmulRSymm2Dxy_param_6];
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r1, %r5, %r6, %r7;
	mov.u32 	%r8, %ntid.y;
	mov.u32 	%r9, %ctaid.y;
	mov.u32 	%r10, %tid.y;
	mad.lo.s32 	%r2, %r8, %r9, %r10;
	setp.ge.s32	%p1, %r2, %r4;
	setp.ge.s32	%p2, %r1, %r3;
	or.pred  	%p3, %p1, %p2;
	@%p3 bra 	BB0_2;

	cvta.to.global.u64 	%rd6, %rd2;
	cvta.to.global.u64 	%rd7, %rd1;
	cvta.to.global.u64 	%rd8, %rd3;
	mad.lo.s32 	%r11, %r2, %r3, %r1;
	shl.b32 	%r12, %r11, 1;
	mul.wide.s32 	%rd9, %r12, 4;
	add.s64 	%rd10, %rd7, %rd9;
	ld.global.f32 	%f1, [%rd10+4];
	add.s64 	%rd11, %rd6, %rd9;
	ld.global.f32 	%f2, [%rd11+4];
	shr.u32 	%r13, %r4, 31;
	add.s32 	%r14, %r4, %r13;
	shr.s32 	%r15, %r14, 1;
	setp.gt.s32	%p4, %r2, %r15;
	sub.s32 	%r16, %r4, %r2;
	selp.b32	%r17, %r16, %r2, %p4;
	selp.f32	%f3, 0fBF800000, 0f3F800000, %p4;
	mad.lo.s32 	%r18, %r17, %r3, %r1;
	mul.wide.s32 	%rd12, %r18, 4;
	add.s64 	%rd13, %rd8, %rd12;
	cvta.to.global.u64 	%rd14, %rd4;
	add.s64 	%rd15, %rd14, %rd12;
	ld.global.nc.f32 	%f4, [%rd15];
	cvta.to.global.u64 	%rd16, %rd5;
	add.s64 	%rd17, %rd16, %rd12;
	ld.global.nc.f32 	%f5, [%rd17];
	mul.f32 	%f6, %f3, %f5;
	ld.global.nc.f32 	%f7, [%rd13];
	ld.global.f32 	%f8, [%rd10];
	ld.global.f32 	%f9, [%rd11];
	mul.f32 	%f10, %f9, %f6;
	fma.rn.f32 	%f11, %f8, %f7, %f10;
	st.global.f32 	[%rd10], %f11;
	mul.f32 	%f12, %f2, %f6;
	fma.rn.f32 	%f13, %f1, %f7, %f12;
	st.global.f32 	[%rd10+4], %f13;
	mul.f32 	%f14, %f8, %f6;
	fma.rn.f32 	%f15, %f9, %f4, %f14;
	st.global.f32 	[%rd11], %f15;
	mul.f32 	%f16, %f1, %f6;
	fma.rn.f32 	%f17, %f2, %f4, %f16;
	st.global.f32 	[%rd11+4], %f17;

BB0_2:
	ret;
}


`
	kernmulRSymm2Dxy_ptx_53 = `
.version 6.5
.target sm_53
.address_size 64

	// .globl	kernmulRSymm2Dxy

.visible .entry kernmulRSymm2Dxy(
	.param .u64 kernmulRSymm2Dxy_param_0,
	.param .u64 kernmulRSymm2Dxy_param_1,
	.param .u64 kernmulRSymm2Dxy_param_2,
	.param .u64 kernmulRSymm2Dxy_param_3,
	.param .u64 kernmulRSymm2Dxy_param_4,
	.param .u32 kernmulRSymm2Dxy_param_5,
	.param .u32 kernmulRSymm2Dxy_param_6
)
{
	.reg .pred 	%p<5>;
	.reg .f32 	%f<18>;
	.reg .b32 	%r<19>;
	.reg .b64 	%rd<18>;


	ld.param.u64 	%rd1, [kernmulRSymm2Dxy_param_0];
	ld.param.u64 	%rd2, [kernmulRSymm2Dxy_param_1];
	ld.param.u64 	%rd3, [kernmulRSymm2Dxy_param_2];
	ld.param.u64 	%rd4, [kernmulRSymm2Dxy_param_3];
	ld.param.u64 	%rd5, [kernmulRSymm2Dxy_param_4];
	ld.param.u32 	%r3, [kernmulRSymm2Dxy_param_5];
	ld.param.u32 	%r4, [kernmulRSymm2Dxy_param_6];
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r1, %r5, %r6, %r7;
	mov.u32 	%r8, %ntid.y;
	mov.u32 	%r9, %ctaid.y;
	mov.u32 	%r10, %tid.y;
	mad.lo.s32 	%r2, %r8, %r9, %r10;
	setp.ge.s32	%p1, %r2, %r4;
	setp.ge.s32	%p2, %r1, %r3;
	or.pred  	%p3, %p1, %p2;
	@%p3 bra 	BB0_2;

	cvta.to.global.u64 	%rd6, %rd2;
	cvta.to.global.u64 	%rd7, %rd1;
	cvta.to.global.u64 	%rd8, %rd3;
	mad.lo.s32 	%r11, %r2, %r3, %r1;
	shl.b32 	%r12, %r11, 1;
	mul.wide.s32 	%rd9, %r12, 4;
	add.s64 	%rd10, %rd7, %rd9;
	ld.global.f32 	%f1, [%rd10+4];
	add.s64 	%rd11, %rd6, %rd9;
	ld.global.f32 	%f2, [%rd11+4];
	shr.u32 	%r13, %r4, 31;
	add.s32 	%r14, %r4, %r13;
	shr.s32 	%r15, %r14, 1;
	setp.gt.s32	%p4, %r2, %r15;
	sub.s32 	%r16, %r4, %r2;
	selp.b32	%r17, %r16, %r2, %p4;
	selp.f32	%f3, 0fBF800000, 0f3F800000, %p4;
	mad.lo.s32 	%r18, %r17, %r3, %r1;
	mul.wide.s32 	%rd12, %r18, 4;
	add.s64 	%rd13, %rd8, %rd12;
	cvta.to.global.u64 	%rd14, %rd4;
	add.s64 	%rd15, %rd14, %rd12;
	ld.global.nc.f32 	%f4, [%rd15];
	cvta.to.global.u64 	%rd16, %rd5;
	add.s64 	%rd17, %rd16, %rd12;
	ld.global.nc.f32 	%f5, [%rd17];
	mul.f32 	%f6, %f3, %f5;
	ld.global.nc.f32 	%f7, [%rd13];
	ld.global.f32 	%f8, [%rd10];
	ld.global.f32 	%f9, [%rd11];
	mul.f32 	%f10, %f9, %f6;
	fma.rn.f32 	%f11, %f8, %f7, %f10;
	st.global.f32 	[%rd10], %f11;
	mul.f32 	%f12, %f2, %f6;
	fma.rn.f32 	%f13, %f1, %f7, %f12;
	st.global.f32 	[%rd10+4], %f13;
	mul.f32 	%f14, %f8, %f6;
	fma.rn.f32 	%f15, %f9, %f4, %f14;
	st.global.f32 	[%rd11], %f15;
	mul.f32 	%f16, %f1, %f6;
	fma.rn.f32 	%f17, %f2, %f4, %f16;
	st.global.f32 	[%rd11+4], %f17;

BB0_2:
	ret;
}


`
	kernmulRSymm2Dxy_ptx_60 = `
.version 6.5
.target sm_60
.address_size 64

	// .globl	kernmulRSymm2Dxy

.visible .entry kernmulRSymm2Dxy(
	.param .u64 kernmulRSymm2Dxy_param_0,
	.param .u64 kernmulRSymm2Dxy_param_1,
	.param .u64 kernmulRSymm2Dxy_param_2,
	.param .u64 kernmulRSymm2Dxy_param_3,
	.param .u64 kernmulRSymm2Dxy_param_4,
	.param .u32 kernmulRSymm2Dxy_param_5,
	.param .u32 kernmulRSymm2Dxy_param_6
)
{
	.reg .pred 	%p<5>;
	.reg .f32 	%f<18>;
	.reg .b32 	%r<19>;
	.reg .b64 	%rd<18>;


	ld.param.u64 	%rd1, [kernmulRSymm2Dxy_param_0];
	ld.param.u64 	%rd2, [kernmulRSymm2Dxy_param_1];
	ld.param.u64 	%rd3, [kernmulRSymm2Dxy_param_2];
	ld.param.u64 	%rd4, [kernmulRSymm2Dxy_param_3];
	ld.param.u64 	%rd5, [kernmulRSymm2Dxy_param_4];
	ld.param.u32 	%r3, [kernmulRSymm2Dxy_param_5];
	ld.param.u32 	%r4, [kernmulRSymm2Dxy_param_6];
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r1, %r5, %r6, %r7;
	mov.u32 	%r8, %ntid.y;
	mov.u32 	%r9, %ctaid.y;
	mov.u32 	%r10, %tid.y;
	mad.lo.s32 	%r2, %r8, %r9, %r10;
	setp.ge.s32	%p1, %r2, %r4;
	setp.ge.s32	%p2, %r1, %r3;
	or.pred  	%p3, %p1, %p2;
	@%p3 bra 	BB0_2;

	cvta.to.global.u64 	%rd6, %rd2;
	cvta.to.global.u64 	%rd7, %rd1;
	cvta.to.global.u64 	%rd8, %rd3;
	mad.lo.s32 	%r11, %r2, %r3, %r1;
	shl.b32 	%r12, %r11, 1;
	mul.wide.s32 	%rd9, %r12, 4;
	add.s64 	%rd10, %rd7, %rd9;
	ld.global.f32 	%f1, [%rd10+4];
	add.s64 	%rd11, %rd6, %rd9;
	ld.global.f32 	%f2, [%rd11+4];
	shr.u32 	%r13, %r4, 31;
	add.s32 	%r14, %r4, %r13;
	shr.s32 	%r15, %r14, 1;
	setp.gt.s32	%p4, %r2, %r15;
	sub.s32 	%r16, %r4, %r2;
	selp.b32	%r17, %r16, %r2, %p4;
	selp.f32	%f3, 0fBF800000, 0f3F800000, %p4;
	mad.lo.s32 	%r18, %r17, %r3, %r1;
	mul.wide.s32 	%rd12, %r18, 4;
	add.s64 	%rd13, %rd8, %rd12;
	cvta.to.global.u64 	%rd14, %rd4;
	add.s64 	%rd15, %rd14, %rd12;
	ld.global.nc.f32 	%f4, [%rd15];
	cvta.to.global.u64 	%rd16, %rd5;
	add.s64 	%rd17, %rd16, %rd12;
	ld.global.nc.f32 	%f5, [%rd17];
	mul.f32 	%f6, %f3, %f5;
	ld.global.nc.f32 	%f7, [%rd13];
	ld.global.f32 	%f8, [%rd10];
	ld.global.f32 	%f9, [%rd11];
	mul.f32 	%f10, %f9, %f6;
	fma.rn.f32 	%f11, %f8, %f7, %f10;
	st.global.f32 	[%rd10], %f11;
	mul.f32 	%f12, %f2, %f6;
	fma.rn.f32 	%f13, %f1, %f7, %f12;
	st.global.f32 	[%rd10+4], %f13;
	mul.f32 	%f14, %f8, %f6;
	fma.rn.f32 	%f15, %f9, %f4, %f14;
	st.global.f32 	[%rd11], %f15;
	mul.f32 	%f16, %f1, %f6;
	fma.rn.f32 	%f17, %f2, %f4, %f16;
	st.global.f32 	[%rd11+4], %f17;

BB0_2:
	ret;
}


`
	kernmulRSymm2Dxy_ptx_61 = `
.version 6.5
.target sm_61
.address_size 64

	// .globl	kernmulRSymm2Dxy

.visible .entry kernmulRSymm2Dxy(
	.param .u64 kernmulRSymm2Dxy_param_0,
	.param .u64 kernmulRSymm2Dxy_param_1,
	.param .u64 kernmulRSymm2Dxy_param_2,
	.param .u64 kernmulRSymm2Dxy_param_3,
	.param .u64 kernmulRSymm2Dxy_param_4,
	.param .u32 kernmulRSymm2Dxy_param_5,
	.param .u32 kernmulRSymm2Dxy_param_6
)
{
	.reg .pred 	%p<5>;
	.reg .f32 	%f<18>;
	.reg .b32 	%r<19>;
	.reg .b64 	%rd<18>;


	ld.param.u64 	%rd1, [kernmulRSymm2Dxy_param_0];
	ld.param.u64 	%rd2, [kernmulRSymm2Dxy_param_1];
	ld.param.u64 	%rd3, [kernmulRSymm2Dxy_param_2];
	ld.param.u64 	%rd4, [kernmulRSymm2Dxy_param_3];
	ld.param.u64 	%rd5, [kernmulRSymm2Dxy_param_4];
	ld.param.u32 	%r3, [kernmulRSymm2Dxy_param_5];
	ld.param.u32 	%r4, [kernmulRSymm2Dxy_param_6];
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r1, %r5, %r6, %r7;
	mov.u32 	%r8, %ntid.y;
	mov.u32 	%r9, %ctaid.y;
	mov.u32 	%r10, %tid.y;
	mad.lo.s32 	%r2, %r8, %r9, %r10;
	setp.ge.s32	%p1, %r2, %r4;
	setp.ge.s32	%p2, %r1, %r3;
	or.pred  	%p3, %p1, %p2;
	@%p3 bra 	BB0_2;

	cvta.to.global.u64 	%rd6, %rd2;
	cvta.to.global.u64 	%rd7, %rd1;
	cvta.to.global.u64 	%rd8, %rd3;
	mad.lo.s32 	%r11, %r2, %r3, %r1;
	shl.b32 	%r12, %r11, 1;
	mul.wide.s32 	%rd9, %r12, 4;
	add.s64 	%rd10, %rd7, %rd9;
	ld.global.f32 	%f1, [%rd10+4];
	add.s64 	%rd11, %rd6, %rd9;
	ld.global.f32 	%f2, [%rd11+4];
	shr.u32 	%r13, %r4, 31;
	add.s32 	%r14, %r4, %r13;
	shr.s32 	%r15, %r14, 1;
	setp.gt.s32	%p4, %r2, %r15;
	sub.s32 	%r16, %r4, %r2;
	selp.b32	%r17, %r16, %r2, %p4;
	selp.f32	%f3, 0fBF800000, 0f3F800000, %p4;
	mad.lo.s32 	%r18, %r17, %r3, %r1;
	mul.wide.s32 	%rd12, %r18, 4;
	add.s64 	%rd13, %rd8, %rd12;
	cvta.to.global.u64 	%rd14, %rd4;
	add.s64 	%rd15, %rd14, %rd12;
	ld.global.nc.f32 	%f4, [%rd15];
	cvta.to.global.u64 	%rd16, %rd5;
	add.s64 	%rd17, %rd16, %rd12;
	ld.global.nc.f32 	%f5, [%rd17];
	mul.f32 	%f6, %f3, %f5;
	ld.global.nc.f32 	%f7, [%rd13];
	ld.global.f32 	%f8, [%rd10];
	ld.global.f32 	%f9, [%rd11];
	mul.f32 	%f10, %f9, %f6;
	fma.rn.f32 	%f11, %f8, %f7, %f10;
	st.global.f32 	[%rd10], %f11;
	mul.f32 	%f12, %f2, %f6;
	fma.rn.f32 	%f13, %f1, %f7, %f12;
	st.global.f32 	[%rd10+4], %f13;
	mul.f32 	%f14, %f8, %f6;
	fma.rn.f32 	%f15, %f9, %f4, %f14;
	st.global.f32 	[%rd11], %f15;
	mul.f32 	%f16, %f1, %f6;
	fma.rn.f32 	%f17, %f2, %f4, %f16;
	st.global.f32 	[%rd11+4], %f17;

BB0_2:
	ret;
}


`
	kernmulRSymm2Dxy_ptx_62 = `
.version 6.5
.target sm_62
.address_size 64

	// .globl	kernmulRSymm2Dxy

.visible .entry kernmulRSymm2Dxy(
	.param .u64 kernmulRSymm2Dxy_param_0,
	.param .u64 kernmulRSymm2Dxy_param_1,
	.param .u64 kernmulRSymm2Dxy_param_2,
	.param .u64 kernmulRSymm2Dxy_param_3,
	.param .u64 kernmulRSymm2Dxy_param_4,
	.param .u32 kernmulRSymm2Dxy_param_5,
	.param .u32 kernmulRSymm2Dxy_param_6
)
{
	.reg .pred 	%p<5>;
	.reg .f32 	%f<18>;
	.reg .b32 	%r<19>;
	.reg .b64 	%rd<18>;


	ld.param.u64 	%rd1, [kernmulRSymm2Dxy_param_0];
	ld.param.u64 	%rd2, [kernmulRSymm2Dxy_param_1];
	ld.param.u64 	%rd3, [kernmulRSymm2Dxy_param_2];
	ld.param.u64 	%rd4, [kernmulRSymm2Dxy_param_3];
	ld.param.u64 	%rd5, [kernmulRSymm2Dxy_param_4];
	ld.param.u32 	%r3, [kernmulRSymm2Dxy_param_5];
	ld.param.u32 	%r4, [kernmulRSymm2Dxy_param_6];
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r1, %r5, %r6, %r7;
	mov.u32 	%r8, %ntid.y;
	mov.u32 	%r9, %ctaid.y;
	mov.u32 	%r10, %tid.y;
	mad.lo.s32 	%r2, %r8, %r9, %r10;
	setp.ge.s32	%p1, %r2, %r4;
	setp.ge.s32	%p2, %r1, %r3;
	or.pred  	%p3, %p1, %p2;
	@%p3 bra 	BB0_2;

	cvta.to.global.u64 	%rd6, %rd2;
	cvta.to.global.u64 	%rd7, %rd1;
	cvta.to.global.u64 	%rd8, %rd3;
	mad.lo.s32 	%r11, %r2, %r3, %r1;
	shl.b32 	%r12, %r11, 1;
	mul.wide.s32 	%rd9, %r12, 4;
	add.s64 	%rd10, %rd7, %rd9;
	ld.global.f32 	%f1, [%rd10+4];
	add.s64 	%rd11, %rd6, %rd9;
	ld.global.f32 	%f2, [%rd11+4];
	shr.u32 	%r13, %r4, 31;
	add.s32 	%r14, %r4, %r13;
	shr.s32 	%r15, %r14, 1;
	setp.gt.s32	%p4, %r2, %r15;
	sub.s32 	%r16, %r4, %r2;
	selp.b32	%r17, %r16, %r2, %p4;
	selp.f32	%f3, 0fBF800000, 0f3F800000, %p4;
	mad.lo.s32 	%r18, %r17, %r3, %r1;
	mul.wide.s32 	%rd12, %r18, 4;
	add.s64 	%rd13, %rd8, %rd12;
	cvta.to.global.u64 	%rd14, %rd4;
	add.s64 	%rd15, %rd14, %rd12;
	ld.global.nc.f32 	%f4, [%rd15];
	cvta.to.global.u64 	%rd16, %rd5;
	add.s64 	%rd17, %rd16, %rd12;
	ld.global.nc.f32 	%f5, [%rd17];
	mul.f32 	%f6, %f3, %f5;
	ld.global.nc.f32 	%f7, [%rd13];
	ld.global.f32 	%f8, [%rd10];
	ld.global.f32 	%f9, [%rd11];
	mul.f32 	%f10, %f9, %f6;
	fma.rn.f32 	%f11, %f8, %f7, %f10;
	st.global.f32 	[%rd10], %f11;
	mul.f32 	%f12, %f2, %f6;
	fma.rn.f32 	%f13, %f1, %f7, %f12;
	st.global.f32 	[%rd10+4], %f13;
	mul.f32 	%f14, %f8, %f6;
	fma.rn.f32 	%f15, %f9, %f4, %f14;
	st.global.f32 	[%rd11], %f15;
	mul.f32 	%f16, %f1, %f6;
	fma.rn.f32 	%f17, %f2, %f4, %f16;
	st.global.f32 	[%rd11+4], %f17;

BB0_2:
	ret;
}


`
	kernmulRSymm2Dxy_ptx_70 = `
.version 6.5
.target sm_70
.address_size 64

	// .globl	kernmulRSymm2Dxy

.visible .entry kernmulRSymm2Dxy(
	.param .u64 kernmulRSymm2Dxy_param_0,
	.param .u64 kernmulRSymm2Dxy_param_1,
	.param .u64 kernmulRSymm2Dxy_param_2,
	.param .u64 kernmulRSymm2Dxy_param_3,
	.param .u64 kernmulRSymm2Dxy_param_4,
	.param .u32 kernmulRSymm2Dxy_param_5,
	.param .u32 kernmulRSymm2Dxy_param_6
)
{
	.reg .pred 	%p<5>;
	.reg .f32 	%f<18>;
	.reg .b32 	%r<19>;
	.reg .b64 	%rd<18>;


	ld.param.u64 	%rd1, [kernmulRSymm2Dxy_param_0];
	ld.param.u64 	%rd2, [kernmulRSymm2Dxy_param_1];
	ld.param.u64 	%rd3, [kernmulRSymm2Dxy_param_2];
	ld.param.u64 	%rd4, [kernmulRSymm2Dxy_param_3];
	ld.param.u64 	%rd5, [kernmulRSymm2Dxy_param_4];
	ld.param.u32 	%r3, [kernmulRSymm2Dxy_param_5];
	ld.param.u32 	%r4, [kernmulRSymm2Dxy_param_6];
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r1, %r5, %r6, %r7;
	mov.u32 	%r8, %ntid.y;
	mov.u32 	%r9, %ctaid.y;
	mov.u32 	%r10, %tid.y;
	mad.lo.s32 	%r2, %r8, %r9, %r10;
	setp.ge.s32	%p1, %r2, %r4;
	setp.ge.s32	%p2, %r1, %r3;
	or.pred  	%p3, %p1, %p2;
	@%p3 bra 	BB0_2;

	cvta.to.global.u64 	%rd6, %rd2;
	cvta.to.global.u64 	%rd7, %rd1;
	cvta.to.global.u64 	%rd8, %rd3;
	mad.lo.s32 	%r11, %r2, %r3, %r1;
	shl.b32 	%r12, %r11, 1;
	mul.wide.s32 	%rd9, %r12, 4;
	add.s64 	%rd10, %rd7, %rd9;
	ld.global.f32 	%f1, [%rd10+4];
	add.s64 	%rd11, %rd6, %rd9;
	ld.global.f32 	%f2, [%rd11+4];
	shr.u32 	%r13, %r4, 31;
	add.s32 	%r14, %r4, %r13;
	shr.s32 	%r15, %r14, 1;
	setp.gt.s32	%p4, %r2, %r15;
	sub.s32 	%r16, %r4, %r2;
	selp.b32	%r17, %r16, %r2, %p4;
	selp.f32	%f3, 0fBF800000, 0f3F800000, %p4;
	mad.lo.s32 	%r18, %r17, %r3, %r1;
	mul.wide.s32 	%rd12, %r18, 4;
	add.s64 	%rd13, %rd8, %rd12;
	cvta.to.global.u64 	%rd14, %rd4;
	add.s64 	%rd15, %rd14, %rd12;
	ld.global.nc.f32 	%f4, [%rd15];
	cvta.to.global.u64 	%rd16, %rd5;
	add.s64 	%rd17, %rd16, %rd12;
	ld.global.nc.f32 	%f5, [%rd17];
	mul.f32 	%f6, %f3, %f5;
	ld.global.nc.f32 	%f7, [%rd13];
	ld.global.f32 	%f8, [%rd10];
	ld.global.f32 	%f9, [%rd11];
	mul.f32 	%f10, %f9, %f6;
	fma.rn.f32 	%f11, %f8, %f7, %f10;
	st.global.f32 	[%rd10], %f11;
	mul.f32 	%f12, %f2, %f6;
	fma.rn.f32 	%f13, %f1, %f7, %f12;
	st.global.f32 	[%rd10+4], %f13;
	mul.f32 	%f14, %f8, %f6;
	fma.rn.f32 	%f15, %f9, %f4, %f14;
	st.global.f32 	[%rd11], %f15;
	mul.f32 	%f16, %f1, %f6;
	fma.rn.f32 	%f17, %f2, %f4, %f16;
	st.global.f32 	[%rd11+4], %f17;

BB0_2:
	ret;
}


`
	kernmulRSymm2Dxy_ptx_72 = `
.version 6.5
.target sm_72
.address_size 64

	// .globl	kernmulRSymm2Dxy

.visible .entry kernmulRSymm2Dxy(
	.param .u64 kernmulRSymm2Dxy_param_0,
	.param .u64 kernmulRSymm2Dxy_param_1,
	.param .u64 kernmulRSymm2Dxy_param_2,
	.param .u64 kernmulRSymm2Dxy_param_3,
	.param .u64 kernmulRSymm2Dxy_param_4,
	.param .u32 kernmulRSymm2Dxy_param_5,
	.param .u32 kernmulRSymm2Dxy_param_6
)
{
	.reg .pred 	%p<5>;
	.reg .f32 	%f<18>;
	.reg .b32 	%r<19>;
	.reg .b64 	%rd<18>;


	ld.param.u64 	%rd1, [kernmulRSymm2Dxy_param_0];
	ld.param.u64 	%rd2, [kernmulRSymm2Dxy_param_1];
	ld.param.u64 	%rd3, [kernmulRSymm2Dxy_param_2];
	ld.param.u64 	%rd4, [kernmulRSymm2Dxy_param_3];
	ld.param.u64 	%rd5, [kernmulRSymm2Dxy_param_4];
	ld.param.u32 	%r3, [kernmulRSymm2Dxy_param_5];
	ld.param.u32 	%r4, [kernmulRSymm2Dxy_param_6];
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r1, %r5, %r6, %r7;
	mov.u32 	%r8, %ntid.y;
	mov.u32 	%r9, %ctaid.y;
	mov.u32 	%r10, %tid.y;
	mad.lo.s32 	%r2, %r8, %r9, %r10;
	setp.ge.s32	%p1, %r2, %r4;
	setp.ge.s32	%p2, %r1, %r3;
	or.pred  	%p3, %p1, %p2;
	@%p3 bra 	BB0_2;

	cvta.to.global.u64 	%rd6, %rd2;
	cvta.to.global.u64 	%rd7, %rd1;
	cvta.to.global.u64 	%rd8, %rd3;
	mad.lo.s32 	%r11, %r2, %r3, %r1;
	shl.b32 	%r12, %r11, 1;
	mul.wide.s32 	%rd9, %r12, 4;
	add.s64 	%rd10, %rd7, %rd9;
	ld.global.f32 	%f1, [%rd10+4];
	add.s64 	%rd11, %rd6, %rd9;
	ld.global.f32 	%f2, [%rd11+4];
	shr.u32 	%r13, %r4, 31;
	add.s32 	%r14, %r4, %r13;
	shr.s32 	%r15, %r14, 1;
	setp.gt.s32	%p4, %r2, %r15;
	sub.s32 	%r16, %r4, %r2;
	selp.b32	%r17, %r16, %r2, %p4;
	selp.f32	%f3, 0fBF800000, 0f3F800000, %p4;
	mad.lo.s32 	%r18, %r17, %r3, %r1;
	mul.wide.s32 	%rd12, %r18, 4;
	add.s64 	%rd13, %rd8, %rd12;
	cvta.to.global.u64 	%rd14, %rd4;
	add.s64 	%rd15, %rd14, %rd12;
	ld.global.nc.f32 	%f4, [%rd15];
	cvta.to.global.u64 	%rd16, %rd5;
	add.s64 	%rd17, %rd16, %rd12;
	ld.global.nc.f32 	%f5, [%rd17];
	mul.f32 	%f6, %f3, %f5;
	ld.global.nc.f32 	%f7, [%rd13];
	ld.global.f32 	%f8, [%rd10];
	ld.global.f32 	%f9, [%rd11];
	mul.f32 	%f10, %f9, %f6;
	fma.rn.f32 	%f11, %f8, %f7, %f10;
	st.global.f32 	[%rd10], %f11;
	mul.f32 	%f12, %f2, %f6;
	fma.rn.f32 	%f13, %f1, %f7, %f12;
	st.global.f32 	[%rd10+4], %f13;
	mul.f32 	%f14, %f8, %f6;
	fma.rn.f32 	%f15, %f9, %f4, %f14;
	st.global.f32 	[%rd11], %f15;
	mul.f32 	%f16, %f1, %f6;
	fma.rn.f32 	%f17, %f2, %f4, %f16;
	st.global.f32 	[%rd11+4], %f17;

BB0_2:
	ret;
}


`
	kernmulRSymm2Dxy_ptx_75 = `
.version 6.5
.target sm_75
.address_size 64

	// .globl	kernmulRSymm2Dxy

.visible .entry kernmulRSymm2Dxy(
	.param .u64 kernmulRSymm2Dxy_param_0,
	.param .u64 kernmulRSymm2Dxy_param_1,
	.param .u64 kernmulRSymm2Dxy_param_2,
	.param .u64 kernmulRSymm2Dxy_param_3,
	.param .u64 kernmulRSymm2Dxy_param_4,
	.param .u32 kernmulRSymm2Dxy_param_5,
	.param .u32 kernmulRSymm2Dxy_param_6
)
{
	.reg .pred 	%p<5>;
	.reg .f32 	%f<18>;
	.reg .b32 	%r<19>;
	.reg .b64 	%rd<18>;


	ld.param.u64 	%rd1, [kernmulRSymm2Dxy_param_0];
	ld.param.u64 	%rd2, [kernmulRSymm2Dxy_param_1];
	ld.param.u64 	%rd3, [kernmulRSymm2Dxy_param_2];
	ld.param.u64 	%rd4, [kernmulRSymm2Dxy_param_3];
	ld.param.u64 	%rd5, [kernmulRSymm2Dxy_param_4];
	ld.param.u32 	%r3, [kernmulRSymm2Dxy_param_5];
	ld.param.u32 	%r4, [kernmulRSymm2Dxy_param_6];
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r1, %r5, %r6, %r7;
	mov.u32 	%r8, %ntid.y;
	mov.u32 	%r9, %ctaid.y;
	mov.u32 	%r10, %tid.y;
	mad.lo.s32 	%r2, %r8, %r9, %r10;
	setp.ge.s32	%p1, %r2, %r4;
	setp.ge.s32	%p2, %r1, %r3;
	or.pred  	%p3, %p1, %p2;
	@%p3 bra 	BB0_2;

	cvta.to.global.u64 	%rd6, %rd2;
	cvta.to.global.u64 	%rd7, %rd1;
	cvta.to.global.u64 	%rd8, %rd3;
	mad.lo.s32 	%r11, %r2, %r3, %r1;
	shl.b32 	%r12, %r11, 1;
	mul.wide.s32 	%rd9, %r12, 4;
	add.s64 	%rd10, %rd7, %rd9;
	ld.global.f32 	%f1, [%rd10+4];
	add.s64 	%rd11, %rd6, %rd9;
	ld.global.f32 	%f2, [%rd11+4];
	shr.u32 	%r13, %r4, 31;
	add.s32 	%r14, %r4, %r13;
	shr.s32 	%r15, %r14, 1;
	setp.gt.s32	%p4, %r2, %r15;
	sub.s32 	%r16, %r4, %r2;
	selp.b32	%r17, %r16, %r2, %p4;
	selp.f32	%f3, 0fBF800000, 0f3F800000, %p4;
	mad.lo.s32 	%r18, %r17, %r3, %r1;
	mul.wide.s32 	%rd12, %r18, 4;
	add.s64 	%rd13, %rd8, %rd12;
	cvta.to.global.u64 	%rd14, %rd4;
	add.s64 	%rd15, %rd14, %rd12;
	ld.global.nc.f32 	%f4, [%rd15];
	cvta.to.global.u64 	%rd16, %rd5;
	add.s64 	%rd17, %rd16, %rd12;
	ld.global.nc.f32 	%f5, [%rd17];
	mul.f32 	%f6, %f3, %f5;
	ld.global.nc.f32 	%f7, [%rd13];
	ld.global.f32 	%f8, [%rd10];
	ld.global.f32 	%f9, [%rd11];
	mul.f32 	%f10, %f9, %f6;
	fma.rn.f32 	%f11, %f8, %f7, %f10;
	st.global.f32 	[%rd10], %f11;
	mul.f32 	%f12, %f2, %f6;
	fma.rn.f32 	%f13, %f1, %f7, %f12;
	st.global.f32 	[%rd10+4], %f13;
	mul.f32 	%f14, %f8, %f6;
	fma.rn.f32 	%f15, %f9, %f4, %f14;
	st.global.f32 	[%rd11], %f15;
	mul.f32 	%f16, %f1, %f6;
	fma.rn.f32 	%f17, %f2, %f4, %f16;
	st.global.f32 	[%rd11+4], %f17;

BB0_2:
	ret;
}


`
)
