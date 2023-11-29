import React, { useState, useEffect } from 'react';
import { Grid, TextField } from '@mui/material';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
import { LocalizationProvider, MobileDateTimePicker } from '@mui/x-date-pickers';
import { DateTimePicker } from '@mui/x-date-pickers/DateTimePicker';
import { useIntl } from 'react-intl';
import { gridSpacing } from 'store/constant';

type PropType = {
    reset: boolean;
    onSelectChange: (value: Date[] | null) => void;
    label: string;
    initialValue?: number;
    viewTime?: boolean | false;
};

const CustomDateTimeRangePicker = React.memo(({ onSelectChange, label, initialValue, reset, viewTime }: PropType) => {
    const intl = useIntl();
    const [startDate, setStartDate] = useState<Date | null>(null);
    const [endDate, setEndDate] = useState<Date | null>(null);

    const handleStartDateChange = (date: Date | null) => {
        setStartDate(date!);
    };

    const handleEndDateChange = (date: Date | null) => {
        if (date) {
            // Set the time to 23:59:59 for end date ONLY as not allowing users to select time
            if (!viewTime) {
                date.setHours(23, 59, 59);
            }
        }
        setEndDate(date);
    };

    const dateTimeFormat = viewTime ? 'yyyy/MM/dd hh:mm a' : 'yyyy/MM/dd';

    useEffect(() => {
        setStartDate(null);
        setEndDate(null);
    }, [reset]);

    useEffect(() => {
        if (startDate && endDate) onSelectChange([startDate, endDate]);
        else onSelectChange(null);
    }, [startDate, endDate]);

    return viewTime ? (
        <LocalizationProvider dateAdapter={AdapterDateFns}>
            <Grid container spacing={gridSpacing}>
                <Grid item xs={12} sm={6}>
                    <MobileDateTimePicker
                        views={['year', 'month', 'day', 'hours', 'minutes']}
                        closeOnSelect={true}
                        componentsProps={{
                            actionBar: {
                                actions: ['today', 'accept', 'cancel', 'clear']
                            }
                        }}
                        inputFormat={dateTimeFormat}
                        label={`${label} ${intl.formatMessage({ id: 'general.start-date' })}`}
                        value={startDate}
                        maxDateTime={endDate ? endDate : undefined}
                        onChange={handleStartDateChange}
                        renderInput={(params: object) => <TextField fullWidth {...params} size="small" />}
                    />
                </Grid>
                <Grid item xs={12} sm={6}>
                    <MobileDateTimePicker
                        views={['year', 'month', 'day', 'hours', 'minutes']}
                        closeOnSelect={true}
                        componentsProps={{
                            actionBar: {
                                actions: ['today', 'accept', 'cancel', 'clear']
                            }
                        }}
                        inputFormat={dateTimeFormat}
                        label={`${label} ${intl.formatMessage({ id: 'general.end-date' })}`}
                        value={endDate}
                        minDateTime={startDate ? startDate : undefined}
                        onChange={handleEndDateChange}
                        renderInput={(params: object) => <TextField fullWidth {...params} size="small" />}
                    />
                </Grid>
            </Grid>
        </LocalizationProvider>
    ) : (
        <LocalizationProvider dateAdapter={AdapterDateFns}>
            <Grid container spacing={gridSpacing}>
                <Grid item xs={12} sm={6}>
                    <DateTimePicker
                        views={['year', 'month', 'day']}
                        closeOnSelect={true}
                        componentsProps={{
                            actionBar: {
                                actions: ['today', 'clear']
                            }
                        }}
                        inputFormat={dateTimeFormat}
                        label={`${label} ${intl.formatMessage({ id: 'general.start-date' })}`}
                        value={startDate}
                        maxDate={endDate ? endDate : undefined}
                        onChange={handleStartDateChange}
                        renderInput={(params: object) => <TextField fullWidth {...params} size="small" />}
                    />
                </Grid>
                <Grid item xs={12} sm={6}>
                    <DateTimePicker
                        views={['year', 'month', 'day']}
                        closeOnSelect={true}
                        componentsProps={{
                            actionBar: {
                                actions: ['today', 'clear']
                            }
                        }}
                        inputFormat={dateTimeFormat}
                        label={`${label} ${intl.formatMessage({ id: 'general.end-date' })}`}
                        value={endDate}
                        minDateTime={startDate ? startDate : undefined}
                        onChange={handleEndDateChange}
                        renderInput={(params: object) => <TextField fullWidth {...params} size="small" />}
                    />
                </Grid>
            </Grid>
        </LocalizationProvider>
    );
});

export default CustomDateTimeRangePicker;
