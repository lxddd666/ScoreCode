export interface ParamForm {
    type?: string;
    name?: string;
    label?: string;
    value?: string | undefined;
    optionButton?: boolean | undefined;
    optionButtonLabel?: string | undefined;
    optionButtonFunc?: () => void;
    optionButtonTimer?: number;
    desc?: string;
    fileChange?: any;
    required?: boolean;
    change?: any;
    InputAdornment?: string;
    emptyButtonLabel?: string;
    options?: any;
    idx1?: string;
    idx2?: string;
    valueKey?: string;
    labelKey?: string;
    style?: { [key: string]: string };
    fileRemove?: any;
    html?: string;
}

export interface FullParamForm {
    formData: ParamForm[];
    formik: any;
    showCancel?: boolean;
    isSubmitting?: boolean;
    customButtonLabel?: string;
    customButtonFunc?: () => void;
    cancelFunc?: () => void;
    dispatch?: any;
    isRequest?: boolean;
    isSubSubmitting?: boolean;
    requiredPermissions?: string[];
}
