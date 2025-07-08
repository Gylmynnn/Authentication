import { fail, message, superValidate } from 'sveltekit-superforms';
import { zod4 } from 'sveltekit-superforms/adapters';
import type { PageServerLoad, Actions } from './$types';
import { loginSchema } from '$lib/schema';
import type { LoginApiResponse } from '$lib/dto';
import { redirect } from '@sveltejs/kit';


export const load: PageServerLoad = async ({ cookies }) => {
    const token = cookies.get('token')
    if (token) {
        redirect(302, '/homepage')
    }
    const form = await superValidate(zod4(loginSchema));
    return { form };
};

export const actions = {
    default: async ({ request, fetch, cookies }) => {
        const form = await superValidate(request, zod4(loginSchema));
        if (!form.valid) {
            return fail(400, { form });
        }
        const res = await fetch('http://localhost:3000/api/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(form.data)
        });

        const responseData: LoginApiResponse = await res.json()
        console.log(responseData)

        if (!res.ok) {
            form.message = responseData.message ?? 'Login gagal.';
            return fail(res.status, { form });
        }

        const { token } = responseData.data;

        cookies.set('token', token, {
            path: '/',
            httpOnly: true,
            sameSite: 'strict',
            secure: false, // ubah jadi false jika local dev tanpa HTTPS
            maxAge: 60 * 60 * 24 // 1 hari
        });

        message(form, 'Login berhasil!');
        return redirect(301, '/homepage')

    }

} satisfies Actions
