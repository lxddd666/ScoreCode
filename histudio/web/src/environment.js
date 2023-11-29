const devEnv = {
    API_URL: 'http://8.222.195.54:4887/',
    AES_KEY: 'f080a463654b2279',
    USERNAME: '',
    PASSWORD: '',
    token: {
        expired: 1 //day
    },
    web: {
        name: 'HiStudio'
    }
}

const stagingEnv = {
    API_URL: 'http://8.222.195.54:4887/',
    AES_KEY: 'f080a463654b2279',
    USERNAME: 'histudio',
    PASSWORD: 'h1Stud10!@#',
    token: {
        expired: 1 //day
    },
    web: {
        name: 'HiStudio'
    }
}

const prodEnv = {
    API_URL: 'http://localhost:8000/',
    AES_KEY: 'f080a463654b2279',
    USERNAME: '',
    PASSWORD: '',
    token: {
        expired: 1 //day
    },
    web: {
        name: 'HiStudio'
    }
}

export default process?.env?.REACT_APP_NODE_ENV == 'production' ? prodEnv : process?.env?.REACT_APP_NODE_ENV == 'staging' ? stagingEnv : devEnv;