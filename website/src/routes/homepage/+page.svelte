<script lang="ts">
    import { enhance } from "$app/forms";
</script>

<span>home</span>
<form action="?/logout" method="POST" use:enhance>
    <button type="submit" name="logout">Logout</button>
</form>
