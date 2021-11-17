// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package plantprovider

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Container struct {
	Brightness *ContainerBrightness `pulumi:"brightness"`
	Color      *string              `pulumi:"color"`
	Material   *string              `pulumi:"material"`
	Size       ContainerSize        `pulumi:"size"`
}

// Defaults sets the appropriate defaults for Container
func (val *Container) Defaults() *Container {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Brightness == nil {
		brightness_ := ContainerBrightness(1.0)
		tmp.Brightness = &brightness_
	}
	return &tmp
}

// ContainerInput is an input type that accepts ContainerArgs and ContainerOutput values.
// You can construct a concrete instance of `ContainerInput` via:
//
//          ContainerArgs{...}
type ContainerInput interface {
	pulumi.Input

	ToContainerOutput() ContainerOutput
	ToContainerOutputWithContext(context.Context) ContainerOutput
}

type ContainerArgs struct {
	Brightness ContainerBrightnessPtrInput `pulumi:"brightness"`
	Color      pulumi.StringPtrInput       `pulumi:"color"`
	Material   pulumi.StringPtrInput       `pulumi:"material"`
	Size       ContainerSizeInput          `pulumi:"size"`
}

func (ContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Container)(nil)).Elem()
}

func (i ContainerArgs) ToContainerOutput() ContainerOutput {
	return i.ToContainerOutputWithContext(context.Background())
}

func (i ContainerArgs) ToContainerOutputWithContext(ctx context.Context) ContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerOutput)
}

func (i ContainerArgs) ToContainerPtrOutput() ContainerPtrOutput {
	return i.ToContainerPtrOutputWithContext(context.Background())
}

func (i ContainerArgs) ToContainerPtrOutputWithContext(ctx context.Context) ContainerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerOutput).ToContainerPtrOutputWithContext(ctx)
}

// ContainerPtrInput is an input type that accepts ContainerArgs, ContainerPtr and ContainerPtrOutput values.
// You can construct a concrete instance of `ContainerPtrInput` via:
//
//          ContainerArgs{...}
//
//  or:
//
//          nil
type ContainerPtrInput interface {
	pulumi.Input

	ToContainerPtrOutput() ContainerPtrOutput
	ToContainerPtrOutputWithContext(context.Context) ContainerPtrOutput
}

type containerPtrType ContainerArgs

func ContainerPtr(v *ContainerArgs) ContainerPtrInput {
	return (*containerPtrType)(v)
}

func (*containerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Container)(nil)).Elem()
}

func (i *containerPtrType) ToContainerPtrOutput() ContainerPtrOutput {
	return i.ToContainerPtrOutputWithContext(context.Background())
}

func (i *containerPtrType) ToContainerPtrOutputWithContext(ctx context.Context) ContainerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerPtrOutput)
}

type ContainerOutput struct{ *pulumi.OutputState }

func (ContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Container)(nil)).Elem()
}

func (o ContainerOutput) ToContainerOutput() ContainerOutput {
	return o
}

func (o ContainerOutput) ToContainerOutputWithContext(ctx context.Context) ContainerOutput {
	return o
}

func (o ContainerOutput) ToContainerPtrOutput() ContainerPtrOutput {
	return o.ToContainerPtrOutputWithContext(context.Background())
}

func (o ContainerOutput) ToContainerPtrOutputWithContext(ctx context.Context) ContainerPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Container) *Container {
		return &v
	}).(ContainerPtrOutput)
}

func (o ContainerOutput) Brightness() ContainerBrightnessPtrOutput {
	return o.ApplyT(func(v Container) *ContainerBrightness { return v.Brightness }).(ContainerBrightnessPtrOutput)
}

func (o ContainerOutput) Color() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Container) *string { return v.Color }).(pulumi.StringPtrOutput)
}

func (o ContainerOutput) Material() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Container) *string { return v.Material }).(pulumi.StringPtrOutput)
}

func (o ContainerOutput) Size() ContainerSizeOutput {
	return o.ApplyT(func(v Container) ContainerSize { return v.Size }).(ContainerSizeOutput)
}

type ContainerPtrOutput struct{ *pulumi.OutputState }

func (ContainerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Container)(nil)).Elem()
}

func (o ContainerPtrOutput) ToContainerPtrOutput() ContainerPtrOutput {
	return o
}

func (o ContainerPtrOutput) ToContainerPtrOutputWithContext(ctx context.Context) ContainerPtrOutput {
	return o
}

func (o ContainerPtrOutput) Elem() ContainerOutput {
	return o.ApplyT(func(v *Container) Container {
		if v != nil {
			return *v
		}
		var ret Container
		return ret
	}).(ContainerOutput)
}

func (o ContainerPtrOutput) Brightness() ContainerBrightnessPtrOutput {
	return o.ApplyT(func(v *Container) *ContainerBrightness {
		if v == nil {
			return nil
		}
		return v.Brightness
	}).(ContainerBrightnessPtrOutput)
}

func (o ContainerPtrOutput) Color() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Container) *string {
		if v == nil {
			return nil
		}
		return v.Color
	}).(pulumi.StringPtrOutput)
}

func (o ContainerPtrOutput) Material() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Container) *string {
		if v == nil {
			return nil
		}
		return v.Material
	}).(pulumi.StringPtrOutput)
}

func (o ContainerPtrOutput) Size() ContainerSizePtrOutput {
	return o.ApplyT(func(v *Container) *ContainerSize {
		if v == nil {
			return nil
		}
		return &v.Size
	}).(ContainerSizePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerInput)(nil)).Elem(), ContainerArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerPtrInput)(nil)).Elem(), ContainerArgs{})
	pulumi.RegisterOutputType(ContainerOutput{})
	pulumi.RegisterOutputType(ContainerPtrOutput{})
}
