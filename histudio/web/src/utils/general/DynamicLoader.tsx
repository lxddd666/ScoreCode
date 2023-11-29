import React, { Suspense, lazy } from 'react';
import Loadable from 'ui-component/Loadable';
import Loader from 'ui-component/Loader';

interface DynamicLoaderProps {
    path: string;
}

const DynamicLoader: React.FC<DynamicLoaderProps> = ({ path }) => {
    let DynamicComponent = Loadable(
        lazy(() =>
            import(`../../views${path}`).then(
                (result) => {
                    return result;
                },
                (error) => {
                    return import('views/pages/maintenance/Error');
                }
            )
        )
    );

    return (
        <Suspense fallback={<Loader />}>
            <DynamicComponent />
        </Suspense>
    );
};

export default DynamicLoader;
