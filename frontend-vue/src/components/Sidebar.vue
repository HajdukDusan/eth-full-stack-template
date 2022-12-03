<template>
    <div>
        <h4 class="sidebar-header pt-3">Sidebar</h4>
        <hr>
        <ul class="nav nav-pills flex-column mb-auto px-3">
            <li v-for="item in items" :key="item.name" class="nav-item">
                <span @click="active = item.name" class="nav-link" :class="{ active: active == item.name }">{{ item.name }}</span>
            </li>
        </ul>
    </div>
</template>

<script>
export default {
    name: 'Sidebar',

    data() {
        return {
            active: "Dashboard",
            items: [
                {
                    name: "Dashboard",
                    to: '/'
                },
                {
                    name: "Wallet",
                    to: '/wallet'
                }
            ]
            
        };
    },

    mounted() {
        
    },

    methods: {
        
    },
};
</script>

<style scoped>

.sidebar-header {
    text-align: center;
}

</style>