import { fail, message, superValidate } from 'sveltekit-superforms';
import { zod4 } from 'sveltekit-superforms/adapters';
import type { PageServerLoad, Actions } from './$types';
import { registerSchema } from '$lib/schema';
import type { RegisterApiResponse } from '$lib/dto';
import { redirect } from '@sveltejs/kit';


export const load: PageServerLoad = async ({ cookies }) => {
    const token = cookies.get('token')
    if (token) {
        redirect(302, '/homepage')
    }
    const form = await superValidate(zod4(registerSchema));
    return { form };
};

export const actions = {
    default: async ({ request, fetch }) => {
        const form = await superValidate(request, zod4(registerSchema));
        if (!form.valid) {
            return fail(400, { form });
        }
        const res = await fetch('http://localhost:3000/api/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(form.data)
        });

        const responseData: RegisterApiResponse = await res.json()
        console.log(responseData)

        if (!res.ok) {
            form.message = responseData.message ?? 'Login gagal.';
            return fail(res.status, { form });
        }

        message(form, 'Register berhasil silahkan login!');
        return redirect(301, '/auth/login')

    }

} satisfies Actions
