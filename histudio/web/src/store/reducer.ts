// third-party
import { combineReducers } from 'redux';
import { persistReducer } from 'redux-persist';
import storage from 'redux-persist/lib/storage';

// project imports
import snackbarReducer from './slices/snackbar';
import userReducer from './slices/user';
import whatsReducer from './slices/whats';
import cronReducer from './slices/cron';
import subcronReducer from './slices/subcron';
import blacklistReducer from './slices/blacklist';
import dictReducer from './slices/dict';
import cartReducer from './slices/cart';
import menuReducer from './slices/menu';
import logReducer from './slices/log';
import loginlogReducer from './slices/loginlog';
import servelogReducer from './slices/servelog';
import smslogReducer from './slices/smslog';
import tgReducer from './slices/tg';
import orgReducer from './slices/org';
import scriptReducer from './slices/script';

// ==============================|| COMBINE REDUCER ||============================== //

const reducer = combineReducers({
    snackbar: snackbarReducer,
    cart: persistReducer(
        {
            key: 'cart',
            storage,
            keyPrefix: 'berry-'
        },
        cartReducer
    ),
    user: userReducer,
    whats: whatsReducer,
    menu: menuReducer,
    log: logReducer,
    loginlog: loginlogReducer,
    servelog: servelogReducer,
    smslog: smslogReducer,
    blacklist: blacklistReducer,
    dict: dictReducer,
    cron: cronReducer,
    subcron: subcronReducer,
    tg: tgReducer,
    org: orgReducer,
    script: scriptReducer
});

export default reducer;
