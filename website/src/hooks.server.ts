import { redirect, type Handle } from '@sveltejs/kit';

export const handle: Handle = async ({ event, resolve }) => {
    const token = event.cookies.get('token');
    const access = event.url.pathname.startsWith("/homepage");
    if (!token && access) {
        throw redirect(303, '/auth/login')
    }
    return resolve(event);
};
