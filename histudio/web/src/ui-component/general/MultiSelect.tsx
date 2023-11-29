import * as React from 'react';
import OutlinedInput from '@mui/material/OutlinedInput';
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import FormControl from '@mui/material/FormControl';
import ListItemText from '@mui/material/ListItemText';
import Select from '@mui/material/Select';
import Checkbox from '@mui/material/Checkbox';
import { FlatOption } from 'types/option';

const ITEM_HEIGHT = 48;
const ITEM_PADDING_TOP = 8;
const MenuProps = {
    PaperProps: {
        style: {
            maxHeight: ITEM_HEIGHT * 4.5 + ITEM_PADDING_TOP
        }
    }
};

type PropType = {
    label: string;
    id: string;
    onSelectChange: (value: number[]) => void;
    options: FlatOption[];
    size?: 'medium' | 'small';
    initialValue?: number[];
};

export default function MultipleSelectCheckmarks({ label, id, options, onSelectChange, initialValue, size = 'medium' }: PropType) {
    const [selectedValue, setSelectedValue] = React.useState<number[]>(initialValue ? initialValue : []);
    const [selectedValueLabel, setSelectedValueLabel] = React.useState<string[]>(initialValue ? handleInitialValueName(initialValue) : []);

    function handleInitialValueName(initialValue: number[]) {
        const valueNames: string[] = [];
        initialValue.map((x) => {
            const selectedOption = options.find((selected) => selected.id === x);
            if (selectedOption) {
                valueNames.push(selectedOption.name);
            }
        });
        return valueNames;
    }

    const handleChange = (value: number, label: string) => {
        if (selectedValue.includes(value)) {
            setSelectedValue(selectedValue.filter((item) => item !== value));
        } else {
            setSelectedValue([...selectedValue, value]);
        }

        if (selectedValueLabel.includes(label)) {
            setSelectedValueLabel(selectedValueLabel.filter((item) => item !== label));
        } else {
            setSelectedValueLabel([...selectedValueLabel, label]);
        }
    };
    React.useEffect(() => {
        onSelectChange(selectedValue);
    }, [selectedValue]);

    return (
        <FormControl fullWidth>
            <InputLabel htmlFor={id}>{label}</InputLabel>
            <Select
                id={id}
                multiple
                value={selectedValueLabel}
                input={<OutlinedInput label={label} />}
                renderValue={(selected) => selected.join(', ')}
                MenuProps={MenuProps}
                className="multiSelect-select"
                size={size}
            >
                {options.map((option) => (
                    <MenuItem
                        key={option.id}
                        value={option.id}
                        onClick={() => {
                            const selectedOption = options.find((selected) => selected.id === option.id);
                            if (selectedOption) {
                                handleChange(selectedOption.id, selectedOption.name);
                            }
                        }}
                    >
                        <Checkbox checked={selectedValueLabel.indexOf(option.name) > -1} />
                        <ListItemText primary={option.name} />
                    </MenuItem>
                ))}
            </Select>
        </FormControl>
    );
}
