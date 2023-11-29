import React, { useState } from 'react';
import TableRow from '@mui/material/TableRow';
import TableCell from '@mui/material/TableCell';
import IconButton from '@mui/material/IconButton';
import AddIcon from '@mui/icons-material/AddCircleOutlineTwoTone';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';
import ToggleOnIcon from '@mui/icons-material/ToggleOnTwoTone';
import ToggleOffIcon from '@mui/icons-material/ToggleOffOutlined';
import ExpandCircleDownIcon from '@mui/icons-material/ExpandCircleDownTwoTone';
import Chip from 'ui-component/extended/Chip';

import { useIntl } from 'react-intl';
import { FieldToShow } from 'types/general';
import { Button, Grid } from '@mui/material';
import { gridSpacing } from 'store/constant';
import { RoleListData } from 'types/user';

interface Props {
    data: any;
    level?: number;
    fieldsToShow?: FieldToShow[];
    /* Props below are for actions column, if not wanted, don't pass in */
    firstRowDisableControl?: boolean;
    handleActivate?: (id: number, status: number) => void;
    handleAddModal?: (id?: number, type?: 'add' | 'edit') => void;
    handleDeleteModal?: (id: number[]) => void;
    handleMenuPermissionModal?: (selectedRole: RoleListData) => void;
    handleDataPermissionModal?: (selectedRole: RoleListData) => void;
}

const ExpandableTableRow = React.memo(
    ({
        data,
        level = 1,
        firstRowDisableControl = false,
        handleActivate,
        handleAddModal,
        handleDeleteModal,
        fieldsToShow,
        handleMenuPermissionModal,
        handleDataPermissionModal
    }: Props) => {
        const intl = useIntl();

        const getAllIdsWithChildren = (data: any): number[] => {
            const result: number[] = [100];
            const traverse = (items: any) => {
                items.forEach((item: any) => {
                    if (item.children && item.children.length > 0) {
                        result.push(item.id);
                        traverse(item.children);
                    }
                });
            };
            traverse(data);
            return result;
        };

        const initExpanded = getAllIdsWithChildren(data);
        const [expandedRows, setExpandedRows] = useState<number[]>(initExpanded);

        const toggleRow = (rowId: number) => {
            if (expandedRows.includes(rowId)) {
                setExpandedRows(expandedRows.filter((id) => id !== rowId));
            } else {
                setExpandedRows([...expandedRows, rowId]);
            }
        };

        React.useEffect(() => {
            setExpandedRows(initExpanded);
        }, [data]);

        // Handling fields to show dynamically, may add on
        // if there are more field types in the future,
        // add onto this function
        // the first column is fixed to row.label
        // while the last column is fixed to actions
        function handleFieldsToShow(row: any) {
            return fieldsToShow!.map((field: FieldToShow, index: number) => {
                switch (field.fieldType) {
                    case 'status':
                        return (
                            <TableCell key={index} align="center">
                                {row[field.fieldName] === 1 && (
                                    <Chip label={intl.formatMessage({ id: 'general.normal' })} size="small" chipcolor="success" />
                                )}
                                {row[field.fieldName] === 2 && (
                                    <Chip label={intl.formatMessage({ id: 'general.disabled' })} size="small" chipcolor="orange" />
                                )}
                            </TableCell>
                        );
                    case 'isDefault':
                        return (
                            <TableCell key={index} align="center">
                                {row[field.fieldName] === 0 && (
                                    <Chip label={intl.formatMessage({ id: 'general.yes' })} size="small" chipcolor="success" />
                                )}
                                {row[field.fieldName] !== 0 && (
                                    <Chip label={intl.formatMessage({ id: 'general.no' })} size="small" chipcolor="error" />
                                )}
                            </TableCell>
                        );
                    case '':
                    default:
                        return (
                            <TableCell key={index} align="center">
                                {row[field.fieldName]}
                            </TableCell>
                        );
                }
            });
        }

        const renderRows = (rows: any) => {
            return rows.map((row: any, index: number) => (
                <React.Fragment key={`${row.id}_${index}`}>
                    <TableRow hover role="checkbox" tabIndex={-1}>
                        <TableCell style={{ paddingLeft: `${level * 16}px` }} align="left">
                            {row.children ? (
                                <IconButton onClick={() => toggleRow(row.id)} size="medium">
                                    <ExpandCircleDownIcon
                                        fontSize="small"
                                        sx={{ transform: `${expandedRows.includes(row.id) ? 'rotate(180deg)' : 'rotate(0deg)'}` }}
                                    />
                                </IconButton>
                            ) : (
                                ''
                            )}
                            {row.label}
                        </TableCell>
                        {fieldsToShow && handleFieldsToShow(row)}
                        <TableCell align="center" className="sticky" sx={{ pr: 3, right: 0 }}>
                            <Grid
                                container
                                spacing={gridSpacing}
                                sx={{ visibility: `${firstRowDisableControl ? 'hidden' : 'visible'}`, justifyContent: 'space-between' }}
                            >
                                {handleMenuPermissionModal && (
                                    <Grid item sm={12} md={6}>
                                        <Button
                                            variant="outlined"
                                            onClick={() => handleMenuPermissionModal(row)}
                                            color="secondary"
                                            size="small"
                                            aria-label="Edit Menu Permission"
                                        >
                                            {intl.formatMessage({ id: 'general.edit-menu-permission' })}
                                        </Button>
                                    </Grid>
                                )}
                                {handleDataPermissionModal && (
                                    <Grid item sm={12} md={6}>
                                        <Button
                                            variant="outlined"
                                            onClick={() => handleDataPermissionModal(row)}
                                            color="secondary"
                                            size="small"
                                            aria-label="Edit Data Permission"
                                        >
                                            {intl.formatMessage({ id: 'general.edit-data-permission' })}
                                        </Button>
                                    </Grid>
                                )}
                                {handleActivate && (
                                    <Grid item xs={6} sm={4} md={3}>
                                        <IconButton
                                            onClick={() => handleActivate(row.id, row.status)}
                                            color="secondary"
                                            size="medium"
                                            aria-label="View"
                                        >
                                            {row.status === 1 ? (
                                                <ToggleOnIcon sx={{ fontSize: '1.3rem' }} />
                                            ) : (
                                                <ToggleOffIcon sx={{ fontSize: '1.3rem' }} />
                                            )}
                                        </IconButton>
                                    </Grid>
                                )}
                                {handleAddModal && (
                                    <Grid item xs={6} sm={4} md={3}>
                                        <IconButton
                                            onClick={() => handleAddModal(row.id, 'add')}
                                            color="secondary"
                                            size="medium"
                                            aria-label="Edit"
                                        >
                                            <AddIcon sx={{ fontSize: '1.3rem' }} />
                                        </IconButton>
                                    </Grid>
                                )}
                                {handleAddModal && (
                                    <Grid item xs={6} sm={4} md={3}>
                                        <IconButton
                                            onClick={() => handleAddModal(row.id, 'edit')}
                                            color="secondary"
                                            size="medium"
                                            aria-label="Edit"
                                        >
                                            <EditIcon sx={{ fontSize: '1.3rem' }} />
                                        </IconButton>
                                    </Grid>
                                )}
                                {handleDeleteModal && (
                                    <Grid item xs={6} sm={4} md={3}>
                                        <IconButton
                                            onClick={() => handleDeleteModal([row.id])}
                                            color="secondary"
                                            size="medium"
                                            aria-label="Delete"
                                        >
                                            <DeleteIcon sx={{ fontSize: '1.3rem' }} />
                                        </IconButton>
                                    </Grid>
                                )}
                            </Grid>
                        </TableCell>
                    </TableRow>
                    {expandedRows.includes(row.id) && row.children && (
                        <ExpandableTableRow
                            data={row.children}
                            level={level + 1}
                            fieldsToShow={fieldsToShow}
                            handleActivate={handleActivate}
                            handleAddModal={handleAddModal}
                            handleDeleteModal={handleDeleteModal}
                            handleMenuPermissionModal={handleMenuPermissionModal}
                            handleDataPermissionModal={handleDataPermissionModal}
                        />
                    )}
                </React.Fragment>
            ));
        };

        return <>{renderRows(data)}</>;
    }
);

export default ExpandableTableRow;
