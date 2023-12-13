// import Box from '@mui/material/Box';
import { memo } from 'react';
import PropTypes from 'prop-types';
// import { WrapperTabs } from './style02.js';

const TabPanel02 = (props: any) => {
    const { value, index, ...other } = props;

    return value === index ? (
        <div
            role="tabpanel"
            hidden={value !== index}
            id={`vertical-tabpanel-${index}`}
            aria-labelledby={`vertical-tab-${index}`}
            {...other}
        >
            {value === index && <div className="tab2_main">1</div>}
        </div>
    ) : (
        <></>
    );
};
TabPanel02.propTypes = {
    children: PropTypes.node,
    index: PropTypes.number.isRequired,
    value: PropTypes.number.isRequired
};
export default memo(TabPanel02);
