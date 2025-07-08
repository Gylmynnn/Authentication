import { z } from "zod/v4";

export const loginSchema = z.object({
    username: z.string().max(16, "Username maksimal 16 karakter").min(3, "Username minimal 3 karakter"),
    password: z.string().max(16, "Password maksimal 16 karakter").min(6, "Password minimal 6 karakter"),
});

export const registerSchema = z.object({
    name: z.string().max(16, "Nama maksimal 16 karakter").min(3, "Nama minimal 3 karakter"),
    username: z.string().max(16, "Username maksimal 16 karakter").min(3, "Username minimal 3 karakter"),
    email: z.email("Email tidak valid example: example@gmail.com"),
    password: z.string().max(16, "Password maksimal 16 karakter").min(6, "Password minimal 6 karakter"),
});
