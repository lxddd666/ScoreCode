import { NestedSubOption, FlatOption } from 'types/option';

function flattenNestedObjects(nestedOptions: NestedSubOption[]): FlatOption[] {
    const flatOptions: FlatOption[] = [];

    function recursiveFlatten(options: NestedSubOption[]) {
        for (const option of options) {
            flatOptions.push({
                id: option.value,
                name: option.name
            });

            if (option.children && option.children.length > 0) {
                recursiveFlatten(option.children);
            }
        }
    }

    recursiveFlatten(nestedOptions);

    return flatOptions;
}

export default flattenNestedObjects;
