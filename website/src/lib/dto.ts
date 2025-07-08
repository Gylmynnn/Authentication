interface LoginUserResponse {
    id: number;
    name: string;
    username: string;
    email: string;
    created_at: string;
    updated_at: string;
    token: string;
}

interface RegisterUserResponse {
    id: number;
    name: string;
    username: string;
    email: string;
    created_at: string;
    updated_at: string;
}

export interface LoginApiResponse {
    success: boolean;
    status: number;
    message: string;
    data: LoginUserResponse
}

export interface RegisterApiResponse {
    success: boolean;
    status: number;
    message: string;
    data: RegisterUserResponse
}
