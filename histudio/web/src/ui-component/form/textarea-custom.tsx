import { TextareaAutosize, TextField } from '@mui/material';
import { useRef } from 'react';

interface TextAreaCustomParam {
    value: string;
    name: string;
    label: string;
    required: boolean;
    changeTextArea: (e: any, data: string) => void;
}

function TextAreaCustom({
    value = '',
    name = '',
    label = '',
    required = false,
    changeTextArea = (e: any, data: string) => {}
}: TextAreaCustomParam) {
    const textAreaRef = useRef<HTMLTextAreaElement | null>(null);

    return (
        <>
            <TextField
                fullWidth
                label={
                    <div>
                        {label}
                        <span style={{ color: 'red' }}>{required ? ' *' : ''}</span>
                    </div>
                }
                InputProps={{
                    inputComponent: TextareaAutosize,
                    inputProps: {
                        minRows: 5,
                        maxRows: 5,
                        ref: textAreaRef,
                        value: value,
                        onChange: (e: any) => {
                            changeTextArea(e, name);
                        },
                        id: `textarea_${name}`,
                        'aria-label': `textarea_${name}`,

                        onFocus: () => {},
                        onBlur: () => {},
                        sx: { padding: 0, margin: '15.5px 14px', borderRadius: 0,  resize: 'vertical'}
                    }
                }}
            />
        </>
    );
}

export default TextAreaCustom;
