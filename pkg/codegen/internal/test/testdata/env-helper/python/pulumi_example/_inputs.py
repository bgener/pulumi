# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'HelmReleaseSettingsArgs',
    'LayeredTypeArgs',
]

@pulumi.input_type
class HelmReleaseSettingsArgs:
    def __init__(__self__, *,
                 required_arg: pulumi.Input[str],
                 driver: Optional[pulumi.Input[str]] = None,
                 plugins_path: Optional[pulumi.Input[str]] = None):
        """
        BETA FEATURE - Options to configure the Helm Release resource.
        :param pulumi.Input[str] required_arg: to test required args
        :param pulumi.Input[str] driver: The backend storage driver for Helm. Values are: configmap, secret, memory, sql.
        :param pulumi.Input[str] plugins_path: The path to the helm plugins directory.
        """
        pulumi.set(__self__, "required_arg", required_arg)
        if driver is None:
            driver = (_utilities.get_env('PULUMI_K8S_HELM_DRIVER') or 'secret')
        if driver is not None:
            pulumi.set(__self__, "driver", driver)
        if plugins_path is None:
            plugins_path = _utilities.get_env('PULUMI_K8S_HELM_PLUGINS_PATH')
        if plugins_path is not None:
            pulumi.set(__self__, "plugins_path", plugins_path)

    @property
    @pulumi.getter(name="requiredArg")
    def required_arg(self) -> pulumi.Input[str]:
        """
        to test required args
        """
        return pulumi.get(self, "required_arg")

    @required_arg.setter
    def required_arg(self, value: pulumi.Input[str]):
        pulumi.set(self, "required_arg", value)

    @property
    @pulumi.getter
    def driver(self) -> Optional[pulumi.Input[str]]:
        """
        The backend storage driver for Helm. Values are: configmap, secret, memory, sql.
        """
        return pulumi.get(self, "driver")

    @driver.setter
    def driver(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "driver", value)

    @property
    @pulumi.getter(name="pluginsPath")
    def plugins_path(self) -> Optional[pulumi.Input[str]]:
        """
        The path to the helm plugins directory.
        """
        return pulumi.get(self, "plugins_path")

    @plugins_path.setter
    def plugins_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "plugins_path", value)


@pulumi.input_type
class LayeredTypeArgs:
    def __init__(__self__, *,
                 other: pulumi.Input['HelmReleaseSettingsArgs'],
                 thinker: pulumi.Input[str],
                 answer: Optional[pulumi.Input[float]] = None,
                 question: Optional[pulumi.Input[str]] = None,
                 recursive: Optional[pulumi.Input['LayeredTypeArgs']] = None):
        """
        Make sure that defaults propagate through types
        :param pulumi.Input[str] thinker: To ask and answer
        :param pulumi.Input[float] answer: The answer to the question
        :param pulumi.Input[str] question: The question already answered
        """
        pulumi.set(__self__, "other", other)
        if thinker is None:
            thinker = 'not a good interaction'
        pulumi.set(__self__, "thinker", thinker)
        if answer is None:
            answer = 42
        if answer is not None:
            pulumi.set(__self__, "answer", answer)
        if question is None:
            question = (_utilities.get_env('PULUMI_THE_QUESTION') or '<unknown>')
        if question is not None:
            pulumi.set(__self__, "question", question)
        if recursive is not None:
            pulumi.set(__self__, "recursive", recursive)

    @property
    @pulumi.getter
    def other(self) -> pulumi.Input['HelmReleaseSettingsArgs']:
        return pulumi.get(self, "other")

    @other.setter
    def other(self, value: pulumi.Input['HelmReleaseSettingsArgs']):
        pulumi.set(self, "other", value)

    @property
    @pulumi.getter
    def thinker(self) -> pulumi.Input[str]:
        """
        To ask and answer
        """
        return pulumi.get(self, "thinker")

    @thinker.setter
    def thinker(self, value: pulumi.Input[str]):
        pulumi.set(self, "thinker", value)

    @property
    @pulumi.getter
    def answer(self) -> Optional[pulumi.Input[float]]:
        """
        The answer to the question
        """
        return pulumi.get(self, "answer")

    @answer.setter
    def answer(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "answer", value)

    @property
    @pulumi.getter
    def question(self) -> Optional[pulumi.Input[str]]:
        """
        The question already answered
        """
        return pulumi.get(self, "question")

    @question.setter
    def question(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "question", value)

    @property
    @pulumi.getter
    def recursive(self) -> Optional[pulumi.Input['LayeredTypeArgs']]:
        return pulumi.get(self, "recursive")

    @recursive.setter
    def recursive(self, value: Optional[pulumi.Input['LayeredTypeArgs']]):
        pulumi.set(self, "recursive", value)


