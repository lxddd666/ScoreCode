import React, { useRef } from 'react';
import { MenuItem, Collapse, IconButton, InputLabel, useTheme, OutlinedInput, Select, FormControl, Radio } from '@mui/material';
import { ArrowDropDown, ArrowRight } from '@mui/icons-material';

type PropType = {
    options: any;
    id: string;
    onSelectChange: (value: number[]) => void;
    size?: 'small' | 'medium';
    label?: string;
    initialValue?: number[];
    disableParent?: boolean;
    multiSelect?: boolean;
    valueFieldName?: string;
    displayFieldName?: string;
    triggerExpandAll?: boolean;
    triggerCheckAll?: boolean;
    filterStatus?: boolean;
};

const findExpandedIds = (options: any[]) => {
    const expandedIds: number[] = [];
    if (options.length > 0)
        options.forEach((option) => {
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

const ExpandableSelect = React.memo(
    ({
        options,
        id,
        onSelectChange,
        size = 'medium',
        initialValue,
        label,
        disableParent = false,
        multiSelect = false,
        displayFieldName = 'label',
        valueFieldName = 'value',
        triggerExpandAll = true,
        triggerCheckAll = false,
        filterStatus = true
    }: PropType) => {
        const theme = useTheme();
        const selectRef = useRef<HTMLSelectElement>(null);
        const filteredOptions = filterStatus ? options.filter((obj: any) => obj.status === 1) : options;
        let selectAllVal: number[] = [];
        const selectAll = React.useMemo(() => {
            function recursiveSelectAll(optionsToSearch: any[]): number[] {
                const values: number[] = [];
                for (const option of optionsToSearch) {
                    values.push(option.id);
                    if (option.children && option.children.length > 0) {
                        const foundValues = recursiveSelectAll(option.children);
                        if (foundValues) {
                            values.push(...foundValues);
                        }
                    }
                }
                return values;
            }
            selectAllVal = recursiveSelectAll(filteredOptions);
            return selectAllVal;
        }, [selectAllVal]);

        const [open, setOpen] = React.useState<boolean>(false);
        const [expandedItems, setExpandedItems] = React.useState<number[]>(triggerExpandAll ? findExpandedIds(filteredOptions) : []);
        const [selectedValue, setSelectedValue] = React.useState<number[]>(triggerCheckAll ? selectAll : initialValue ? initialValue : []);
        const [selectedValueLabel, setSelectedValueLabel] = React.useState<string[]>(findValueName(selectedValue));

        function findValueName(value: number[]) {
            function recursiveFindValueName(optionsToSearch: any[]): string[] {
                const ValueNames: string[] = [];
                for (const option of optionsToSearch) {
                    if (value.includes(option[valueFieldName])) {
                        ValueNames.push(option[displayFieldName]);
                    }
                    if (option.children && option.children.length > 0) {
                        const foundName = recursiveFindValueName(option.children);
                        if (foundName) {
                            ValueNames.push(...foundName);
                        }
                    }
                }
                return ValueNames;
            }
            return recursiveFindValueName(filteredOptions);
        }

        const handleExpandToggle = (id: number) => {
            if (expandedItems.includes(id)) {
                setExpandedItems(expandedItems.filter((item) => item !== id));
            } else {
                setExpandedItems([...expandedItems, id]);
            }
        };

        function isNumberArrayEqual(arr1: number[], arr2: number[]): boolean {
            if (arr1.length !== arr2.length) {
                return false;
            }

            for (let i = 0; i < arr1.length; i++) {
                if (arr1[i] !== arr2[i]) {
                    return false;
                }
            }

            return true;
        }

        React.useEffect(() => {
            // Adding this checking to avoid infinite looping
            if (initialValue && !isNumberArrayEqual(initialValue, selectedValue)) {
                setSelectedValue(initialValue);
            }
        }, [initialValue]);

        React.useEffect(() => {
            // Adding this checking to avoid infinite looping
            if (!isNumberArrayEqual(initialValue || [], selectedValue)) {
                onSelectChange(selectedValue);
            }
            setSelectedValueLabel(findValueName(selectedValue));
        }, [selectedValue]);

        // Forcing the code below to render on 2nd render and above ONLY
        const firstRender = React.useRef(true);
        React.useEffect(() => {
            if (firstRender.current) {
                firstRender.current = false;
                return;
            } else {
                if (triggerCheckAll) {
                    const selected = selectAll;
                    setSelectedValue(selected);
                    setSelectedValueLabel(findValueName(selected));
                } else {
                    setSelectedValue([]);
                    setSelectedValueLabel([]);
                }
            }
        }, [triggerCheckAll]);

        React.useEffect(() => {
            if (triggerExpandAll) {
                setExpandedItems(selectAll);
            } else {
                setExpandedItems([]);
            }
        }, [triggerExpandAll]);

        const renderOptions = (filteredOptions: any[], level = 0) => {
            return filteredOptions.map((option) => (
                <div key={option[valueFieldName]}>
                    <MenuItem
                        value={option[valueFieldName]}
                        onClick={(e) => {
                            const selectedOption = filteredOptions.find((selected) => selected[valueFieldName] === option[valueFieldName]);
                            if (selectedOption) {
                                if (multiSelect) {
                                    if (selectedValue.includes(selectedOption[valueFieldName])) {
                                        setSelectedValue((prev) => prev.filter((val) => val !== selectedOption[valueFieldName]));
                                        setSelectedValueLabel((prev) => prev.filter((str) => str !== selectedOption[displayFieldName]));
                                    } else {
                                        setSelectedValue((prev) => [...prev, selectedOption[valueFieldName]]);
                                        setSelectedValueLabel((prev) => [...prev, selectedOption[displayFieldName]]);
                                    }
                                } else {
                                    setSelectedValue([selectedOption[valueFieldName]]);
                                    setSelectedValueLabel([selectedOption[displayFieldName]]);
                                }
                            }
                            if (!multiSelect) setOpen(false);
                        }}
                        style={{ paddingLeft: `${level * 16}px` }}
                        disabled={disableParent && option.children}
                    >
                        {!disableParent && (
                            <IconButton
                                onClick={(e) => {
                                    handleExpandToggle(option.id);
                                    e.stopPropagation();
                                }}
                                size="small"
                                sx={{
                                    color: theme.palette.secondary.main,
                                    visibility: `${option.children ? (option.children[0].status === 1 ? 'visible' : 'hidden') : 'hidden'}`
                                }}
                            >
                                {expandedItems.includes(option.id) ? <ArrowDropDown /> : <ArrowRight />}
                            </IconButton>
                        )}
                        {multiSelect && (
                            <Radio size="small" value={option[valueFieldName]} checked={selectedValue.includes(option[valueFieldName])} />
                        )}
                        {option[displayFieldName]}
                    </MenuItem>
                    <Collapse in={expandedItems.includes(option.id)}>
                        {option.children && renderOptions(option.children, level + 1)}
                    </Collapse>
                </div>
            ));
        };

        return (
            <FormControl fullWidth className="ExpandableSelect">
                {label && (
                    <InputLabel sx={{ backgroundColor: 'inherit' }} style={{ transform: 'translate(14px, 5px) scale(0.75)' }} htmlFor={id}>
                        {label}
                    </InputLabel>
                )}
                <Select
                    open={open}
                    ref={selectRef}
                    id={id}
                    onClose={() => setOpen(false)}
                    onOpen={() => setOpen(true)}
                    renderValue={(value) => value.join(', ')}
                    value={selectedValueLabel}
                    input={<OutlinedInput label={selectedValueLabel} sx={{ '& div': { padding: '30.5px 14px 11.5px!important' } }} />}
                    multiple={multiSelect}
                >
                    {renderOptions(filteredOptions)}
                </Select>
            </FormControl>
        );
    }
);

export default ExpandableSelect;
