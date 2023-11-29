import { useMemo, ReactNode } from 'react';

// material-ui
import { CssBaseline, StyledEngineProvider } from '@mui/material';
import { createTheme, ThemeOptions, ThemeProvider, Theme, TypographyVariantsOptions } from '@mui/material/styles';

// project import
import useConfig from 'hooks/useConfig';
import Palette from './palette';
import Typography from './typography';

import componentStyleOverrides from './compStyleOverride';
import customShadows from './shadows';

// types
import { CustomShadowProps } from 'types/default-theme';
import React from 'react';

interface Props {
    children: ReactNode;
}

export default function ThemeCustomization({ children }: Props) {
    const { borderRadius, fontFamily, navType, outlinedFilled, presetColor, rtlLayout } = useConfig();

    const theme: Theme = useMemo<Theme>(() => Palette(navType, presetColor), [navType, presetColor]);

    // setting up var to be used in css
    React.useEffect(() => {
        document.documentElement.style.setProperty('--white', theme.palette.mode === 'dark' ? '#030614' : '#ffffff');
        document.documentElement.style.setProperty('--gray', theme.palette.mode === 'dark' ? '#333333' : '#eeeeee');
        document.documentElement.style.setProperty(
            '--primary-main',
            theme.palette.mode === 'dark' ? theme.palette.primary.dark : theme.palette.primary.light
        );
        document.documentElement.style.setProperty(
            '--secondary-main',
            theme.palette.mode === 'dark' ? theme.palette.secondary.dark : theme.palette.secondary.light
        );
        document.documentElement.style.setProperty(
            '--error-main',
            theme.palette.mode === 'dark' ? theme.palette.error.dark : theme.palette.error.light
        );
        document.documentElement.style.setProperty(
            '--warning-main',
            theme.palette.mode === 'dark' ? theme.palette.warning.dark : theme.palette.warning.light
        );
        document.documentElement.style.setProperty(
            '--info-main',
            theme.palette.mode === 'dark' ? theme.palette.info.dark : theme.palette.info.light
        );
        document.documentElement.style.setProperty(
            '--success-main',
            theme.palette.mode === 'dark' ? theme.palette.success.dark : theme.palette.success.light
        );
    }, [theme]);

    // eslint-disable-next-line react-hooks/exhaustive-deps
    const themeTypography: TypographyVariantsOptions = useMemo<TypographyVariantsOptions>(
        () => Typography(theme, borderRadius, fontFamily),
        [theme, borderRadius, fontFamily]
    );
    const themeCustomShadows: CustomShadowProps = useMemo<CustomShadowProps>(() => customShadows(navType, theme), [navType, theme]);

    const themeOptions: ThemeOptions = useMemo(
        () => ({
            direction: rtlLayout ? 'rtl' : 'ltr',
            palette: theme.palette,
            mixins: {
                toolbar: {
                    minHeight: '48px',
                    padding: '16px',
                    '@media (min-width: 600px)': {
                        minHeight: '48px'
                    }
                }
            },
            typography: themeTypography,
            customShadows: themeCustomShadows
        }),
        [rtlLayout, theme, themeCustomShadows, themeTypography]
    );

    const themes: Theme = createTheme(themeOptions);
    themes.components = useMemo(
        () => componentStyleOverrides(themes, borderRadius, outlinedFilled),
        [themes, borderRadius, outlinedFilled]
    );

    return (
        <StyledEngineProvider injectFirst>
            <ThemeProvider theme={themes}>
                <CssBaseline />
                {children}
            </ThemeProvider>
        </StyledEngineProvider>
    );
}
