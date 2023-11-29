// material-ui
import { styled } from '@mui/material/styles';
import { Spin } from 'antd';

// styles
const LoaderWrapper = styled('div')({
    top: 0,
    left: 0,
    zIndex: 1100,
    bottom: 0,
    right: 0
});

const LoaderWrapperChild = styled('div')({
    display:'flex', 
    justifyContent:'center', 
    alignItems: 'center',
    width:'100%', 
    height: '100%'
});

// ==============================|| LOADER ||============================== //

interface loadingType {
    isTransprent? : boolean;
    isFixed? : boolean;
    zIndex?: number
}
const Loading = ({isTransprent, isFixed = true, zIndex = 1099} : loadingType) => (
    <>
        <LoaderWrapper style={{background: isTransprent ? "rgba(0,0,0,0.5)" : "#ccc", position: isFixed ? "fixed" : "absolute", zIndex: zIndex}}>
            <LoaderWrapperChild>
                <Spin size="large" />
            </LoaderWrapperChild>
        </LoaderWrapper>
    </>
);

export default Loading;


