<script setup lang="ts"></script>

<template>
  <div class="container mx-auto my-5 p-5">
    <div class="md:flex no-wrap md:-mx-2">
      <Profile />
      <!-- Right Side -->
      <div class="w-full md:w-9/12 mx-2 h-max border rounded-lg test"></div>
    </div>
  </div>
</template>
<style>
.test {
  min-height: 100vh;
  justify-content: center;
  align-items: center;
  text-align: center;
}
</style>
