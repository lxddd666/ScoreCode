import React from 'react';
// import { Select } from 'antd';
import Icon from '@ant-design/icons';
// import * as icons from '@ant-design/icons';
import { FormControl, useTheme, Select, MenuItem, OutlinedInput, InputLabel, Typography } from '@mui/material';
import { useIntl } from 'react-intl';
const icons = require(`@ant-design/icons`);
export interface IconSelectProps {
    id?: string;
    label?: string;
    onSelectChange: (value: string) => void;
    initialValue: string;
}

const IconSelect: React.FC<IconSelectProps> = ({ id = 'antdPicker', label, onSelectChange, initialValue = '' }) => {
    const iconList = Object.keys(icons)
        .filter((item) => typeof icons[item] === 'object')
        .filter((item) => item === 'default');
    const theme = useTheme();
    const intl = useIntl();

    const [selectedValue, setSelectedValue] = React.useState<string>(initialValue);

    React.useEffect(() => {
        if (initialValue !== selectedValue) setSelectedValue(initialValue);
    }, [initialValue]);

    React.useEffect(() => {
        if (initialValue !== selectedValue) onSelectChange(selectedValue);
    }, [selectedValue]);

    return (
        <FormControl fullWidth className="ExpandableSelect" sx={{ ...theme.typography.customInput }}>
            {label && (
                <InputLabel sx={{ backgroundColor: 'inherit' }} htmlFor={id}>
                    {label}
                </InputLabel>
            )}
            <Select
                id={id}
                input={<OutlinedInput sx={{ '& div': { padding: '30.5px 14px 11.5px!important' } }} />}
                style={{ width: '100%' }}
                value={selectedValue}
            >
                <MenuItem onClick={() => setSelectedValue('')} value={''}>
                    <Typography variant="h6">{intl.formatMessage({ id: 'menu-perm.select-menu-icon' })}</Typography>
                </MenuItem>
                {iconList.map((item) => (
                    <MenuItem onClick={() => setSelectedValue(item)} value={item} key={item}>
                        <Icon
                            component={icons[item]}
                            color="primary"
                            style={{
                                marginRight: '8px'
                            }}
                        />
                        {item}
                    </MenuItem>
                ))}
            </Select>
        </FormControl>
    );
};

export default IconSelect;
