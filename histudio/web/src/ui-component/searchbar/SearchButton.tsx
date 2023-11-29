//material-ui
import FindInPageTwoToneIcon from '@mui/icons-material/FindInPageTwoTone';
import { useIntl } from 'react-intl';
import {
    IconButton
    , Tooltip
} from '@mui/material';

type PropType = {
    onClick?: () => void
}

const SearchButton = ({ onClick }: PropType) => {
    return (
        <Tooltip title={useIntl().formatMessage({ id: "searchbar.search" })}>
            <IconButton size="large" onClick={onClick}>
                <FindInPageTwoToneIcon />
            </IconButton>
        </Tooltip>
    )
}

export default SearchButton;