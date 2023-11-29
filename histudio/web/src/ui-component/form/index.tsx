import {
    Grid,
    Stack,
    Button,
    TextField,
    RadioGroup,
    FormControlLabel,
    Radio,
    FormHelperText,
    InputLabel,
    InputAdornment,
    Typography
} from '@mui/material';
import AnimateButton from 'ui-component/extended/AnimateButton';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
import { LocalizationProvider, DatePicker } from '@mui/x-date-pickers';
import FormControl from '@mui/material/FormControl';
import FormLabel from '@mui/material/FormLabel';
import { Spin } from 'antd';
import { LoadingOutlined } from '@ant-design/icons';
import { FullParamForm } from 'utils/generalInterface';
import { useIntl } from 'react-intl';
import moment from 'moment';
import Switch, { SwitchProps } from '@mui/material/Switch';
import { styled } from '@mui/material/styles';
import { useTheme } from '@mui/material/styles';
import MenuItemGeneral from './menu-item';
import 'react-draft-wysiwyg/dist/react-draft-wysiwyg.css';
import MyEditor from './quill-editor';
import UploadElement from './upload';
import PasswordField from './password';
import TextAreaCustom from './textarea-custom';
import DynamicPermissionComponent from 'ui-component/general/DynamicPermissionComponent';

const antIcon = (
    <LoadingOutlined
        style={{
            fontSize: 20,
            marginLeft: 8
        }}
        spin
    />
);

function GeneralForm({
    formData,
    formik,
    showCancel = false,
    cancelFunc = () => {},
    isSubmitting = false,
    customButtonLabel = '',
    customButtonFunc = () => {},
    dispatch,
    isRequest = false,
    isSubSubmitting = false,
    requiredPermissions = []
}: FullParamForm) {
    const intl = useIntl();
    const themeGeneral = useTheme();

    const SwitchMui = styled((props: SwitchProps) => <Switch focusVisibleClassName=".Mui-focusVisible" disableRipple {...props} />)(
        ({ theme }) => ({
            width: 42,
            height: 26,
            padding: 0,
            '& .MuiSwitch-switchBase': {
                padding: 0,
                margin: 2,
                transitionDuration: '300ms',
                '&.Mui-checked': {
                    transform: 'translateX(16px)',
                    color: '#fff',
                    '& + .MuiSwitch-track': {
                        backgroundColor: themeGeneral?.palette?.secondary?.main,
                        opacity: 1,
                        border: 0
                    },
                    '&.Mui-disabled + .MuiSwitch-track': {
                        opacity: 0.5
                    }
                },
                '&.Mui-focusVisible .MuiSwitch-thumb': {
                    color: '#33cf4d',
                    border: '6px solid #fff'
                },
                '&.Mui-disabled .MuiSwitch-thumb': {
                    color: theme.palette.mode === 'light' ? theme.palette.grey[100] : theme.palette.grey[600]
                },
                '&.Mui-disabled + .MuiSwitch-track': {
                    opacity: theme.palette.mode === 'light' ? 0.7 : 0.3
                }
            },
            '& .MuiSwitch-thumb': {
                boxSizing: 'border-box',
                width: 22,
                height: 22
            },
            '& .MuiSwitch-track': {
                borderRadius: 26 / 2,
                backgroundColor: theme.palette.mode === 'light' ? '#E9E9EA' : '#39393D',
                opacity: 1,
                transition: theme.transitions.create(['background-color'], {
                    duration: 500
                })
            }
        })
    );

    const onChangeGeneral = (value: any, idx: any, data: any, e: any) => {
        if (typeof data?.onChange !== 'undefined') {
            data.onChange(value);
        }
        formik?.setFieldValue(idx, value);
    };

    const onChangeMultiSelect = (value: any, idx: any, data: any, e: any) => {
        if (typeof data?.onChange !== 'undefined') {
            data.onChange(value);
        }

        const tempMultiSelected = formik?.getFieldProps(idx)?.value;

        if (!tempMultiSelected) {
            let newTempMultiSelected = [];
            newTempMultiSelected.push(value);
            formik?.setFieldValue(idx, newTempMultiSelected);
        } else if (Array.isArray(tempMultiSelected)) {
            const findIndex = tempMultiSelected?.findIndex((x) => x == value);
            if (findIndex > -1) {
                tempMultiSelected.splice(findIndex, 1);
            } else {
                tempMultiSelected.push(value);
            }

            formik?.setFieldValue(idx, tempMultiSelected);
        }
    };

    const changeDynamicAddRemoveField = (e: any, keyParent: string, idx: number, key: string) => {
        const dataList = formik?.getFieldProps(keyParent)?.value;
        if (dataList && dataList?.length > 0 && dataList[idx] && key in dataList[idx]) {
            dataList[idx][key] = e?.target?.value;
            formik?.setFieldValue(keyParent, dataList);
        }
    };

    const DynamicAddField = (keyParent: string, idx: number, idx1: string, idx2: string) => {
        const dataList = formik?.getFieldProps(keyParent)?.value;
        if (dataList && dataList?.length > 0) {
            const newAddJson: any = {};
            newAddJson[idx1] = '';
            newAddJson[idx2] = '';
            dataList?.splice(idx + 1, 0, newAddJson);
            formik?.setFieldValue(keyParent, dataList);
        }
    };

    const DynamicRemoveField = (keyParent: string, idx: number) => {
        const dataList = formik?.getFieldProps(keyParent)?.value;
        if (dataList && dataList?.length > 0 && dataList[idx]) {
            dataList?.splice(idx, 1);
            formik?.setFieldValue(keyParent, dataList);
        }
    };

    const addNewDynamicField = (keyParent: string, idx1: string, idx2: string) => {
        const newArrayTemplate = [];
        const newAddJson: any = {};
        newAddJson[idx1] = '';
        newAddJson[idx2] = '';
        newArrayTemplate.push(newAddJson);
        formik?.setFieldValue(keyParent, newArrayTemplate);
    };

    const changeTextArea = (e: any, key: string) => {
        formik?.setFieldValue(key, e?.target?.value);
    };

    const generalFormField = (data: any) => {
        switch (data?.type) {
            case 'radio':
                return (
                    <>
                        <FormControl>
                            <FormLabel>
                                {
                                    <div>
                                        {data?.label}
                                        <span style={{ color: 'red' }}>{data?.required ? ' *' : ''}</span>
                                    </div>
                                }
                            </FormLabel>
                            <RadioGroup
                                name={data?.name}
                                value={formik.values[data?.name]}
                                onChange={formik.handleChange}
                                className="general-custom-radio"
                            >
                                {data?.options?.length > 0 &&
                                    data?.options?.map((item: any, idx: number) => {
                                        return (
                                            <FormControlLabel
                                                key={`${data?.name}_${item?.value}`}
                                                value={item?.value}
                                                control={<Radio />}
                                                label={item?.label}
                                            />
                                        );
                                    })}
                            </RadioGroup>
                        </FormControl>
                        {data?.desc ? <FormHelperText className="input-hint">{data?.desc}</FormHelperText> : undefined}
                        {formik.touched[data?.name] && Boolean(formik.errors[data?.name]) && (
                            <FormHelperText error> {formik.errors[data?.name]} </FormHelperText>
                        )}
                    </>
                );
            case 'number':
                return (
                    <>
                        <TextField
                            fullWidth
                            id={data?.name}
                            name={data?.name}
                            label={
                                <div>
                                    {data?.label}
                                    <span style={{ color: 'red' }}>{data?.required ? ' *' : ''}</span>
                                </div>
                            }
                            value={formik.values[data?.name]}
                            onChange={formik.handleChange}
                            error={formik.touched[data?.name] && Boolean(formik.errors[data?.name])}
                            helperText={data?.desc}
                            type="number"
                            InputProps={{
                                endAdornment: data?.InputAdornment ? (
                                    <InputAdornment position="end">
                                        <InputLabel>{data?.InputAdornment}</InputLabel>
                                    </InputAdornment>
                                ) : undefined
                            }}
                        />
                        {formik.touched[data?.name] && formik.errors[data?.name] ? (
                            <FormHelperText error>{formik.touched[data?.name] && formik.errors[data?.name]}</FormHelperText>
                        ) : undefined}
                    </>
                );
            case 'select':
                return (
                    <>
                        <MenuItemGeneral
                            options={data?.options}
                            id={data?.name}
                            dataAll={data}
                            onSelectChange={onChangeGeneral}
                            value={formik.values[data?.name]}
                            label={data?.label}
                            error={formik.touched[data?.name] && Boolean(formik.errors[data?.name])}
                            helperText={data?.desc}
                            required={data?.required}
                        />
                        {formik.touched[data?.name] && Boolean(formik.errors[data?.name]) && (
                            <FormHelperText error>{formik.touched[data?.name] && formik.errors[data?.name]}</FormHelperText>
                        )}
                    </>
                );
            case 'multiselect':
                return (
                    <>
                        <MenuItemGeneral
                            options={data?.options}
                            id={data?.name}
                            dataAll={data}
                            onSelectChange={onChangeMultiSelect}
                            value={formik.values[data?.name]}
                            label={data?.label}
                            error={formik.touched[data?.name] && Boolean(formik.errors[data?.name])}
                            helperText={data?.desc}
                            isMultiple={true}
                            labelKey={data?.labelKey}
                            valueKey={data?.valueKey}
                        />
                        {formik.touched[data?.name] && Boolean(formik.errors[data?.name]) && (
                            <FormHelperText error>{formik.touched[data?.name] && formik.errors[data?.name]}</FormHelperText>
                        )}
                    </>
                );
            case 'datetime':
                return (
                    <LocalizationProvider dateAdapter={AdapterDateFns}>
                        <DatePicker
                            onChange={(value) =>
                                formik.setFieldValue(data?.name, value ? moment(value)?.format('YYYY-MM-DD HH:mm:ss') : '', true)
                            }
                            value={formik.values[data?.name]}
                            label={
                                <div>
                                    {data?.label}
                                    <span style={{ color: 'red' }}>{data?.required ? ' *' : ''}</span>
                                </div>
                            }
                            renderInput={(params) => (
                                <TextField
                                    label={data?.label}
                                    name={data?.name}
                                    fullWidth
                                    {...params}
                                    error={formik.touched[data?.name] && Boolean(formik.errors[data?.name])}
                                    helperText={data?.desc}
                                />
                            )}
                            inputFormat="dd/MM/yyyy"
                        />
                        {formik.touched[data?.name] && Boolean(formik.errors[data?.name]) && (
                            <FormHelperText error>{formik.touched[data?.name] && formik.errors[data?.name]}</FormHelperText>
                        )}
                    </LocalizationProvider>
                );
            case 'textarea':
                return (
                    <>
                        <FormControl fullWidth>
                            <TextAreaCustom
                                value={formik.values[data?.name]}
                                changeTextArea={(e: any) => changeTextArea(e, data?.name)}
                                name={data?.name}
                                required={data?.required}
                                label={data?.label}
                            />
                        </FormControl>

                        {data?.desc ? <FormHelperText className="input-hint">{data?.desc}</FormHelperText> : undefined}
                        {formik.touched[data?.name] && formik.errors[data?.name] ? (
                            <FormHelperText error>{formik.touched[data?.name] && formik.errors[data?.name]}</FormHelperText>
                        ) : undefined}
                    </>
                );
            case 'upload':
                const serviceToken = window.localStorage.getItem('serviceToken');
                return (
                    <FormControl>
                        <FormLabel>
                            {
                                <div>
                                    {data?.label}
                                    <span style={{ color: 'red' }}>{data?.required ? ' *' : ''}</span>
                                </div>
                            }
                        </FormLabel>
                        <div>
                            <UploadElement
                                onChange={data?.fileChange}
                                fileRemove={data?.fileRemove}
                                headers={{ Authorization: `${serviceToken}` }}
                                data={{ type: '0' }}
                                value={formik.values[data?.name]}
                            />
                        </div>
                        {formik.touched[data?.name] && Boolean(formik.errors[data?.name]) && (
                            <FormHelperText error>{formik.touched[data?.name] && formik.errors[data?.name]}</FormHelperText>
                        )}
                    </FormControl>
                );
            case 'password':
                return (
                    <>
                        <PasswordField
                            name={data?.name}
                            required={data?.required}
                            value={formik.values[data?.name]}
                            handleChange={formik.handleChange}
                            error={formik.touched[data?.name] && Boolean(formik.errors[data?.name])}
                            desc={data?.desc}
                            label={data?.label}
                        />

                        {formik.touched[data?.name] && Boolean(formik.errors[data?.name]) && (
                            <FormHelperText error>{formik.touched[data?.name] && formik.errors[data?.name]}</FormHelperText>
                        )}
                    </>
                );
            case 'textLabel':
                return <div style={{ color: '#aaa' }}>{data?.label}</div>;
            case 'switch':
                return (
                    <div className="switch-container">
                        <FormControlLabel
                            control={
                                <SwitchMui
                                    id={data?.name}
                                    name={data?.name}
                                    checked={formik?.values && formik.values[data?.name] ? formik.values[data?.name] : false}
                                    value={formik.values[data?.name]}
                                    onChange={(eData) =>
                                        typeof data?.change !== 'undefined' ? data?.change(eData?.target?.value) : formik.handleChange
                                    }
                                />
                            }
                            label={
                                <div>
                                    {data?.label}
                                    <span style={{ color: 'red' }}>{data?.required ? ' *' : ''}</span>
                                </div>
                            }
                            labelPlacement="top"
                            className="switch-general-input"
                        />
                        {data?.desc ? <FormHelperText className="input-hint">{data?.desc}</FormHelperText> : undefined}
                        {formik.touched[data?.name] && Boolean(formik.errors[data?.name]) && (
                            <FormHelperText error>{formik.touched[data?.name] && formik.errors[data?.name]}</FormHelperText>
                        )}
                    </div>
                );
            case 'horizontalLineText':
                return (
                    <>
                        <div className="horizontal-line">
                            <hr className="hr-short" />
                            <Typography variant="body1" component="p" sx={{ fontSize: '1.15rem' }}>
                                {data?.label}
                            </Typography>

                            <hr className="hr-long" />
                        </div>
                    </>
                );
            case 'DynamicAddRemoveField':
                if (formik?.getFieldProps(data?.name)?.value?.length > 0) {
                    const outputInputField = formik?.getFieldProps(data?.name)?.value?.map((items: any, index: number) => {
                        if (items && data?.idx1 in items && data?.idx2 in items) {
                            return (
                                <div key={`DynamicAddRemoveField${index}`} className="dynamic-add-remove-field">
                                    <TextField
                                        fullWidth
                                        value={items[data?.idx1]}
                                        onChange={(e) => {
                                            changeDynamicAddRemoveField(e, data?.name, index, data?.idx1);
                                        }}
                                        sx={{ marginRight: 2 }}
                                    />

                                    <TextField
                                        fullWidth
                                        value={items[data?.idx2]}
                                        onChange={(e) => {
                                            changeDynamicAddRemoveField(e, data?.name, index, data?.idx2);
                                        }}
                                        sx={{ marginRight: 2 }}
                                    />
                                    <div style={{ margin: 'auto' }}>
                                        <Button
                                            sx={{ marginRight: 2 }}
                                            variant="contained"
                                            color="secondary"
                                            onClick={() => DynamicRemoveField(data?.name, index)}
                                        >
                                            -
                                        </Button>
                                    </div>
                                    <div style={{ margin: 'auto' }}>
                                        <Button
                                            variant="contained"
                                            color="secondary"
                                            sx={{ height: 'fit-content' }}
                                            onClick={() => DynamicAddField(data?.name, index, data?.idx1, data?.idx2)}
                                        >
                                            +
                                        </Button>
                                    </div>
                                </div>
                            );
                        } else {
                            return <></>;
                        }
                    });

                    return (
                        <>
                            <FormControl fullWidth>
                                <FormLabel>
                                    {
                                        <div>
                                            {data?.label}
                                            <span style={{ color: 'red' }}>{data?.required ? ' *' : ''}</span>
                                        </div>
                                    }
                                </FormLabel>
                                {outputInputField}
                                {formik.touched[data?.name] && Boolean(formik.errors[data?.name]) && (
                                    <FormHelperText error> {formik.errors[data?.name]} </FormHelperText>
                                )}
                            </FormControl>
                        </>
                    );
                } else {
                    return (
                        <FormControl fullWidth>
                            <FormLabel>
                                {
                                    <div>
                                        {data?.label}
                                        <span style={{ color: 'red' }}>{data?.required ? ' *' : ''}</span>
                                    </div>
                                }
                            </FormLabel>
                            <Button
                                color="secondary"
                                variant="contained"
                                onClick={() => addNewDynamicField(data?.name, data?.idx1, data?.idx2)}
                            >
                                {data?.emptyButtonLabel}
                            </Button>
                            {formik.touched[data?.name] && Boolean(formik.errors[data?.name]) && (
                                <FormHelperText error> {formik.errors[data?.name]} </FormHelperText>
                            )}
                        </FormControl>
                    );
                }
            case 'quillEditor':
                return (
                    <div style={{ maxWidth: '100%', position: 'relative' }}>
                        <FormLabel>
                            {
                                <div>
                                    {data?.label}
                                    <span style={{ color: 'red' }}>{data?.required ? ' *' : ''}</span>
                                </div>
                            }
                        </FormLabel>
                        <MyEditor value={formik.values[data?.name]} formik={formik} keyId={data?.name} dispatch={dispatch} />
                        {formik.touched[data?.name] && Boolean(formik.errors[data?.name]) && (
                            <FormHelperText error> {formik.errors[data?.name]} </FormHelperText>
                        )}
                    </div>
                );
            case 'htmlcontent':
                return (
                    <>
                        {data?.label && (
                            <FormLabel>
                                {
                                    <div>
                                        {data?.label}
                                        <span style={{ color: 'red' }}>{data?.required ? ' *' : ''}</span>
                                    </div>
                                }
                            </FormLabel>
                        )}
                        <Typography variant="body1" component="p">
                            <div
                                style={{
                                    ...data?.style,
                                    ...{
                                        color: themeGeneral.palette.secondary.main,
                                        backgroundColor:
                                            themeGeneral.palette.mode === 'dark' ? themeGeneral.palette.background.default : 'transparent',
                                        border:
                                            themeGeneral.palette.mode === 'dark' ? 'none' : `1px solid ${themeGeneral.palette.secondary.main}`
                                    }
                                }}
                                dangerouslySetInnerHTML={{ __html: data?.value }}
                            />
                        </Typography>

                        {data?.desc ? <FormHelperText className="input-hint">{data?.desc}</FormHelperText> : undefined}
                        {formik.touched[data?.name] && Boolean(formik.errors[data?.name]) && (
                            <FormHelperText error> {formik.errors[data?.name]} </FormHelperText>
                        )}
                    </>
                );
            default:
                return (
                    <>
                        <TextField
                            fullWidth
                            id={data?.name}
                            name={data?.name}
                            label={
                                <div>
                                    {data?.label}
                                    <span style={{ color: 'red' }}>{data?.required ? ' *' : ''}</span>
                                </div>
                            }
                            value={formik.values[data?.name]}
                            onChange={formik.handleChange}
                            error={formik.touched[data?.name] && Boolean(formik.errors[data?.name])}
                            helperText={data?.desc}
                            InputProps={{
                                endAdornment: data?.InputAdornment ? (
                                    <InputAdornment position="end">
                                        <InputLabel>{data?.InputAdornment}</InputLabel>
                                    </InputAdornment>
                                ) : undefined
                            }}
                        />
                        {formik.touched[data?.name] && formik.errors[data?.name] ? (
                            <FormHelperText error>{formik.touched[data?.name] && formik.errors[data?.name]}</FormHelperText>
                        ) : undefined}
                    </>
                );
        }
    };

    return (
        <div className="general-form-container">
            <form onSubmit={formik.handleSubmit}>
                <Grid container spacing={2}>
                    {formData &&
                        formData?.length > 0 &&
                        formData?.map((item: any, idx: number) => {
                            if (item?.optionButton) {
                                return (
                                    <Grid className="input-field-container" item xs={12} key={item?.name}>
                                        <div className="input-field-custom">{generalFormField(item)}</div>
                                        <div className="input-field-custom-button">
                                            <Button
                                                disabled={isRequest ? true : false}
                                                onClick={
                                                    typeof item?.optionButtonFunc !== 'undefined' &&
                                                    ((typeof item?.optionButtonTimer == 'number' && item?.optionButtonTimer <= 0) ||
                                                        !item?.optionButtonTimer)
                                                        ? item?.optionButtonFunc
                                                        : () => {}
                                                }
                                                variant="contained"
                                                color="secondary"
                                                style={{ display: 'flex' }}
                                            >
                                                {item?.optionButtonLabel}
                                                {typeof item?.optionButtonTimer == 'number' && item?.optionButtonTimer > 0
                                                    ? ` (${item?.optionButtonTimer})`
                                                    : undefined}
                                                {isRequest ? <Spin indicator={antIcon} /> : <></>}
                                            </Button>
                                        </div>
                                    </Grid>
                                );
                            } else {
                                return (
                                    <Grid item xs={12} key={item?.name}>
                                        {generalFormField(item)}
                                    </Grid>
                                );
                            }
                        })}

                    <Grid item xs={12}>
                        <Stack direction="row" spacing={1} justifyContent="flex-end">
                            {showCancel ? (
                                <AnimateButton>
                                    <Button variant="contained" style={{ display: 'flex' }} color="primary" onClick={cancelFunc}>
                                        {intl.formatMessage({ id: 'general.cancel' })}
                                    </Button>
                                </AnimateButton>
                            ) : (
                                <></>
                            )}

                            {customButtonLabel ? (
                                <AnimateButton>
                                    <Button
                                        variant="contained"
                                        disabled={isSubSubmitting ?? true}
                                        style={{ opacity: isSubSubmitting ? 0.5 : 1, display: 'flex' }}
                                        color="primary"
                                        onClick={customButtonFunc}
                                    >
                                        {customButtonLabel}
                                        {isSubSubmitting ? <Spin indicator={antIcon} /> : <></>}
                                    </Button>
                                </AnimateButton>
                            ) : (
                                <></>
                            )}

                            <DynamicPermissionComponent
                                requiredPermissions={requiredPermissions}
                                children={
                                    <AnimateButton>
                                        <Button
                                            variant="contained"
                                            type="submit"
                                            disabled={isSubmitting ?? true}
                                            style={{ opacity: isSubmitting ? 0.5 : 1, display: 'flex' }}
                                            color="secondary"
                                        >
                                            {intl.formatMessage({ id: 'general.submit' })}
                                            {isSubmitting ? <Spin indicator={antIcon} /> : <></>}
                                        </Button>
                                    </AnimateButton>
                                }
                            />
                        </Stack>
                    </Grid>
                </Grid>
            </form>
        </div>
    );
}

export default GeneralForm;
