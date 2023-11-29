export type HandleFunction = (i: string, s: string) => Promise<void>;

export type BlacklistProfile = {
    id?: number;
    ip?: string;
    remark?: string;
    status?: number;
    createdAt?: string;
    updatedAt?: string;
};

