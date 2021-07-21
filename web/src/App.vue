<template>
    <div id="app">
        <home/>
    </div>
</template>

<script>
    import HelloWorld from './components/HelloWorld.vue'
    import Home from "./views/Home";

    export default {
        name: 'app',
        components: {
            Home,
            HelloWorld
        }
    }
</script>

<style>
    #app {
        font-family: 'Avenir', Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        text-align: center;
        color: #2c3e50;
    }
</style>
