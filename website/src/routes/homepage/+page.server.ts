import type { Actions, PageServerLoad } from "./$types";

// export const load: PageServerLoad = ({ cookies }) => {
//
//     const logout = () => {
//         cookies.delete("token", { path: "/" })
//     }
//
//     return {
//         logout: logout()
//     }
//
// }


export const actions = {
    logout: async ({ cookies }) => {
        cookies.delete('token', { path: '/' });
    }
} satisfies Actions;
