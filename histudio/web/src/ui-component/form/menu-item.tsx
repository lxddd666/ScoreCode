import { useEffect } from 'react';
import useState from 'react-usestateref';
import { MenuItem, Collapse, IconButton, useTheme, FormControl, FormHelperText, TextField } from '@mui/material';
import { ArrowDropDown, ArrowRight } from '@mui/icons-material';
import { NestedSubOption } from 'types/option';
import Checkbox from '@mui/material/Checkbox';

type PropType = {
    options: NestedSubOption[];
    id?: string;
    onSelectChange: (value: number, key?: string, func?: any, e?: any) => void;
    size?: 'small' | 'medium';
    label: string;
    value?: number;
    valueKey?: string;
    labelKey?: string;
    dataAll?: any;
    error?: boolean;
    helperText?: string;
    isMultiple?: boolean;
    required?: boolean;
};

const findExpandedIds = (options: NestedSubOption[]) => {
    const expandedIds: number[] = [];

    options?.length > 0 &&
        options?.forEach((option) => {
            if (option.children) {
                expandedIds.push(option.id);
            }

            if (option.children && Array.isArray(option.children)) {
                const subExpandedIds = findExpandedIds(option.children);
                expandedIds.push(...subExpandedIds);
            }
        });

    return expandedIds;
};

const MenuItemGeneral = ({
    options,
    id,
    onSelectChange,
    size = 'medium',
    value,
    label,
    valueKey = 'value',
    labelKey = 'label',
    dataAll = {},
    error = false,
    helperText,
    isMultiple = false,
    required = false
}: PropType) => {
    const theme = useTheme();
    const [expandedItems, setExpandedItems] = useState<number[]>(findExpandedIds(options));
    const [labelDisplay, setLabelDisplay, labelDisplayRef] = useState<string | string[]>('');
    const [isShow, setIsShow] = useState<boolean>(false);

    useEffect(() => {
        if (isMultiple) {
            if (Array.isArray(value) && value.length > 0) {
                value?.map((item: number, idx: number) => {
                    if (item) {
                        const name = findValueName(item);
                        if (!labelDisplayRef?.current) {
                            let tempLabelDisplay = [];
                            tempLabelDisplay.push(name);
                            setLabelDisplay(tempLabelDisplay);
                        } else if (Array.isArray(labelDisplayRef?.current)) {
                            labelDisplayRef?.current.push(name);
                            setLabelDisplay([...labelDisplayRef?.current]);
                        }
                    }
                });
            }
        } else {
            const name = findValueName(value);
            setLabelDisplay(name);
        }
    }, [value]);

    function findValueName(initialValue?: number) {
        function recursiveFindValueName(optionsToSearch: NestedSubOption[]): string {
            if (optionsToSearch) {
                for (const option of optionsToSearch) {
                    if (option[valueKey]?.toString() === initialValue?.toString()) {
                        return option[labelKey];
                    } else if (option.children && option.children.length > 0) {
                        const foundName = recursiveFindValueName(option.children);
                        if (foundName) {
                            return foundName;
                        }
                    }
                }
            }
            return '';
        }
        return recursiveFindValueName(options);
    }

    const handleExpandToggle = (id: number) => {
        if (expandedItems.includes(id)) {
            setExpandedItems(expandedItems.filter((item) => item !== id));
        } else {
            setExpandedItems([...expandedItems, id]);
        }
    };

    const openSelect = () => {
        setIsShow(true);
    };

    const closeSelect = () => {
        setIsShow(false);
    };

    const renderOptions = (options: NestedSubOption[], level = 0) => {
        return Array.isArray(options) && options?.length > 0 ? (
            options?.map((option) => [
                <MenuItem
                    value={option[valueKey]}
                    onClick={(e) => {
                        if (option[valueKey] !== value) {
                            onSelectChange(option[valueKey], id, dataAll, e);
                            if (!isMultiple) {
                                setIsShow(false);
                            } else {
                                if (Array.isArray(value) && value.length > 0) {
                                    setLabelDisplay([]);
                                    value?.map((item: number, idx: number) => {
                                        if (item) {
                                            const name = findValueName(item);
                                            if (!labelDisplayRef?.current) {
                                                let tempLabelDisplay = [];
                                                tempLabelDisplay.push(name);
                                                setLabelDisplay(tempLabelDisplay);
                                            } else if (Array.isArray(labelDisplayRef?.current)) {
                                                labelDisplayRef?.current.push(name);
                                                setLabelDisplay([...labelDisplayRef?.current]);
                                            }
                                        }
                                    });
                                }
                            }
                        }
                    }}
                    style={{ paddingLeft: `${level * 16}px` }}
                >
                    {option.children && (
                        <IconButton
                            onClick={(e) => {
                                handleExpandToggle(option.id);
                                e.stopPropagation();
                            }}
                            size="small"
                            sx={{ color: theme.palette.secondary.main }}
                        >
                            {expandedItems.includes(option.id) ? <ArrowDropDown /> : <ArrowRight />}
                        </IconButton>
                    )}
                    {isMultiple && <Checkbox checked={Array.isArray(value) ? value?.indexOf(option[valueKey]) > -1 : false} />}
                    {option[labelKey]}
                </MenuItem>,

                <Collapse in={expandedItems.includes(option.id)}>{option.children && renderOptions(option.children, level + 1)}</Collapse>
            ])
        ) : (
            <MenuItem></MenuItem>
        );
    };

    return (
        <>
            <FormControl fullWidth className="ExpandableSelect">
                <TextField
                    select
                    label={
                        <div>
                            <span style={{ color: error ? '#d9534f' : '' }}>{label}</span>
                            <span style={{ color: 'red' }}>{required ? ' *' : ''}</span>
                        </div>
                    }
                    id={id}
                    SelectProps={{
                        open: isShow,
                        renderValue: () =>
                            isMultiple && Array.isArray(labelDisplay) && labelDisplay?.length > 0
                                ? labelDisplay?.join(',')
                                : labelDisplay
                                ? labelDisplay
                                : '',
                        value: value,
                        error: error,
                        onOpen: openSelect,
                        onClose: closeSelect,
                        multiple: isMultiple
                    }}
                >
                    {renderOptions(options)}
                </TextField>

                {helperText && <FormHelperText>{helperText}</FormHelperText>}
                {helperText && <FormHelperText error>{error}</FormHelperText>}
            </FormControl>
        </>
    );
};

export default MenuItemGeneral;
