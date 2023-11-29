import React from 'react';
import { useIntl } from 'react-intl';
import { FlatOption } from 'types/option';
import { FormControl, FormControlLabel, FormLabel, Grid, Radio, RadioGroup, Typography, useTheme } from '@mui/material';

type PropType = {
    options: FlatOption[];
    id: string;
    onSelectChange: (value: number) => void;
    label: string;
    valueFieldName?: 'id' | 'value';
    initialValue?: number;
};

function RadioButtonGroup({
    options,
    id,
    onSelectChange,
    label,
    valueFieldName = 'id',
    initialValue = options[0][valueFieldName]!
}: PropType) {
    const intl = useIntl();
    const theme = useTheme();
    const [selectedValue, setSelectedValue] = React.useState<number>(initialValue);

    const handleChange = (value: number) => {
        setSelectedValue(value);
    };

    React.useEffect(() => {
        // Adding this checking to avoid infinite looping
        if (initialValue !== undefined && initialValue !== selectedValue) {
            setSelectedValue(initialValue);
        }
    }, [initialValue]);

    React.useEffect(() => {
        // Adding this checking to avoid infinite looping
        if (initialValue !== selectedValue) {
            onSelectChange(selectedValue);
        }
    }, [selectedValue]);

    return (
        <FormControl component="fieldset" fullWidth>
            <FormLabel component="legend">{label}</FormLabel>
            <RadioGroup id={id} sx={{ display: 'flex', flexDirection: 'row' }} defaultValue={selectedValue}>
                <Grid container>
                    {options.map((option, index: any) => {
                        return (
                            <Grid
                                key={index}
                                item
                                sm={12}
                                md={6}
                                lg={4}
                                sx={{
                                    width: '100%',
                                    [theme.breakpoints.up('lg')]: {
                                        ':first-of-type': {
                                            borderRadius: '4px',
                                            borderTopRightRadius: '0px',
                                            borderBottomRightRadius: '0px'
                                        },
                                        ':last-of-type': {
                                            borderRadius: '4px',
                                            borderTopLeftRadius: '0px',
                                            borderBottomLeftRadius: '0px'
                                        },
                                        border: '1px solid #ccc',
                                        borderRadius: '0px'
                                    },
                                    border: '1px solid #ccc',
                                    borderRadius: '4px'
                                }}
                            >
                                <FormControl fullWidth>
                                    <FormControlLabel
                                        key={option.id}
                                        value={option[valueFieldName]}
                                        control={
                                            <Radio
                                                size="medium"
                                                className="NoPrefixRadio"
                                                icon={
                                                    <Typography
                                                        sx={{
                                                            display: 'none',
                                                            ':checked': {
                                                                animation: 'unset'
                                                            }
                                                        }}
                                                    />
                                                }
                                                checkedIcon={
                                                    <Typography
                                                        sx={{
                                                            display: 'none',
                                                            ':checked': {
                                                                animation: 'unset'
                                                            }
                                                        }}
                                                    />
                                                }
                                                sx={{
                                                    '& span': {
                                                        display: 'none'
                                                    },
                                                    '&.input': {
                                                        display: 'none'
                                                    }
                                                }}
                                            />
                                        }
                                        onClick={() => {
                                            handleChange(option[valueFieldName]!);
                                        }}
                                        label={<Typography>{intl.formatMessage({ id: option.name })}</Typography>}
                                        sx={{
                                            margin: 0,
                                            paddingY: '0.5rem',
                                            paddingRight: ' 1rem',
                                            cursor: 'pointer',
                                            transition: 'background-color 0.3s',
                                            backgroundColor: 'transparent',
                                            '&:hover': {
                                                color: theme.palette.secondary.main
                                            },
                                            '&.Mui-checked': {
                                                backgroundColor: theme.palette.secondary.light,
                                                color: theme.palette.secondary.main
                                            }
                                        }}
                                        className={`${selectedValue === option[valueFieldName] ? 'Mui-checked' : ''}`}
                                    />
                                </FormControl>
                            </Grid>
                        );
                    })}
                </Grid>
            </RadioGroup>
        </FormControl>
    );
}

export default RadioButtonGroup;
