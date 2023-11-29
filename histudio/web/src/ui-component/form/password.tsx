import { useState } from 'react';
import { TextField } from '@mui/material';
import InputAdornment from '@mui/material/InputAdornment';
import IconButton from '@mui/material/IconButton';
import Visibility from '@mui/icons-material/Visibility';
import VisibilityOff from '@mui/icons-material/VisibilityOff';

interface PasswordFieldParam {
    name?: string;
    value?: string;
    label?: string;
    error?: boolean;
    desc?: string;
    handleChange?: () => void,
    required?: boolean
}
function PasswordField({name = "", value = "", label = "", error = false, desc = "", handleChange = () => {}, required = false} : PasswordFieldParam) {
    const [isShow, setIsShow] = useState<boolean>(false);

    const handlePasswordVisibility = () => {
        setIsShow(!isShow)
    }
    return (
        <TextField
            type={isShow ? 'text' : 'password'}
            fullWidth
            id={name}
            name={name}
            label={
                <div>
                    {label}
                    <span style={{ color: 'red' }}>{required ? ' *' : ''}</span>
                </div>
            }
            value={value}
            onChange={typeof handleChange !== "undefined" ? handleChange : ()=>{}}
            error={error}
            helperText={desc}
            InputProps={{
                endAdornment: (
                  <InputAdornment position="end">
                    <IconButton onClick={handlePasswordVisibility} edge="end">
                      {isShow ? <VisibilityOff /> : <Visibility />}
                    </IconButton>
                  </InputAdornment>
                ),
              }}
        />
    );
}

export default PasswordField;
