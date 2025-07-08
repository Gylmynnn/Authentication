<script lang="ts">
    import { superForm } from "sveltekit-superforms";
    import { zod4 } from "sveltekit-superforms/adapters";
    import { loginSchema } from "$lib/schema";
    import { goto } from "$app/navigation";

    let { data } = $props();
    const { form, errors, message, constraints, enhance } = superForm(
        data.form,
        {
            validators: zod4(loginSchema),
        },
    );
</script>

<main class="h-screen lg:flex lg:justify-center lg:mx-auto">
    <span class="w-full lg:w-1/2">
        <img
            class="p-4 rounded-[32px] lg:h-full object-cover"
            src="/pict.jpeg"
            alt="err"
        />
    </span>

    {#if $message}<h3>{$message}</h3>{/if}

    <div class="lg:flex lg:flex-col lg:justify-center">
        <span class="quotes flex flex-col px-6">
            <h1 class="text-2xl font-bold">Welcome Back</h1>
            <p class="text-[0.8rem] py-4 lg:w-[80%]">
                Dengan satu langkah mudah, kamu bisa kembali ke semua hal yang
                penting dan menyenangkan
            </p>
        </span>

        <div class="px-6 lg:w-[80%]">
            <form method="POST" use:enhance>
                <span class="flex flex-col gap-1">
                    <label class="font-bold" for="name">Name</label>
                    <input
                        placeholder="Masukkan nama"
                        class="bg-blush-black placeholder:text-sm rounded-sm py-3 px-4 text-blush-foreground"
                        type="text"
                        name="name"
                        aria-invalid={$errors.name ? "true" : undefined}
                        bind:value={$form.name}
                        {...$constraints.name}
                    />
                </span>
                {#if $errors.name}<span class="invalid">{$errors.name}</span
                    >{/if}

                <span class="flex flex-col gap-1 pt-2">
                    <label class="font-bold" for="username">Username</label>
                    <input
                        placeholder="Masukkan username"
                        class="bg-blush-black placeholder:text-sm rounded-sm py-3 px-4 text-blush-foreground"
                        type="text"
                        name="username"
                        aria-invalid={$errors.username ? "true" : undefined}
                        bind:value={$form.username}
                        {...$constraints.username}
                    />
                </span>
                {#if $errors.username}<span class="invalid"
                        >{$errors.username}</span
                    >{/if}

                <span class="flex flex-col gap-1 pt-2">
                    <label class="font-bold" for="email">Email</label>
                    <input
                        placeholder="Masukkan email"
                        class="bg-blush-black placeholder:text-sm rounded-sm py-3 px-4 text-blush-foreground"
                        type="email"
                        name="email"
                        aria-invalid={$errors.email ? "true" : undefined}
                        bind:value={$form.email}
                        {...$constraints.email}
                    />
                </span>

                {#if $errors.email}<span class="invalid">{$errors.email}</span
                    >{/if}

                <span class="flex flex-col gap-1 pt-2">
                    <label class="font-bold" for="password">Password</label>
                    <input
                        placeholder="Masukkan password"
                        class="bg-blush-black placeholder:text-sm rounded-sm py-3 px-4 text-blush-foreground"
                        type="password"
                        name="password"
                        aria-invalid={$errors.password ? "true" : undefined}
                        bind:value={$form.password}
                        {...$constraints.username}
                    />
                </span>

                {#if $errors.password}<span class="invalid"
                        >{$errors.password}</span
                    >{/if}

                <div class="flex justify-end pr-2 cursor-pointer">
                    <button
                        type="button"
                        class="text-sm text-blush-blue cursor-pointer"
                        >Generate password!</button
                    >
                </div>

                <div class="flex flex-col py-2">
                    <button
                        class="bg-blush-blue w-full rounded-sm py-2 font-bold"
                        type="submit">Register</button
                    >
                </div>
            </form>

            <div class="flex justify-center gap-1">
                <p class="text-sm text-blush-blue">Sudah punya akun?</p>
                <button
                    onclick={() => goto("/auth/login")}
                    class="text-sm text-blush-blue cursor-pointer">Login</button
                >
            </div>
        </div>
    </div>
</main>

<style>
    .invalid {
        color: red;
    }
</style>
