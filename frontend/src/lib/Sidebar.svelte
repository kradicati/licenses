<script>
    import {Nav, NavItem, NavLink} from "sveltestrap";
</script>

<Nav vertical>
    <NavItem>
        <NavLink href="/licenses">Dashboard</NavLink>
    </NavItem>
    <NavItem>
        <NavLink href="/licenses/list">Licenses</NavLink>
    </NavItem>
    <NavItem>
        <NavLink href="#">Another Link</NavLink>
    </NavItem>
    <NavItem>
        <NavLink disabled href="#">Disabled Link</NavLink>
    </NavItem>
</Nav>