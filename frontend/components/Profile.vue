<template>
  <div class="w-full md:w-3/12 md:mx-2">
    <!-- Profile Card -->
    <div class="bg-white p-3 border border-t-4 border-t-green-400 rounded-lg">
      <img
        class="h-auto w-full mx-auto"
        src="~/assets/images/kamoshika.png"
        alt=""
      />

      <h1 class="text-gray-900 font-bold text-xl leading-8 my-1">Daichi</h1>
      <p class="text-sm text-gray-500 hover:text-gray-600 leading-6">
        新卒で地方銀行に入社し、約3年半勤務。その後、エンジニアとしてキャリアチェンジし、これまでRuby・TypeScript・Goなどを使用した開発業務を行なっています。
      </p>
      <ul
        class="bg-gray-100 text-gray-600 hover:text-gray-700 hover:shadow py-2 px-3 mt-3 divide-y rounded shadow-sm"
      >
        <li class="flex items-center py-3">
          <span>Status</span>
          <span class="ml-auto"
            ><span class="bg-green-500 py-1 px-2 rounded text-white text-sm"
              >Looking for new job</span
            ></span
          >
        </li>
        <li class="flex items-center py-3">
          <span>Birthday</span>
          <span class="ml-auto">Oct 02, 1994</span>
        </li>
      </ul>
    </div>
    <!-- End of profile card -->
    <div class="my-4"></div>
    <!-- Friends card -->
    <div class="bg-white p-3 border rounded-lg">
      <div
        class="flex items-center space-x-3 font-semibold text-gray-900 text-xl leading-8"
      >
        <span class="text-green-500">
          <svg
            class="h-5 fill-current"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"
            />
          </svg>
        </span>
        <span>最新記事</span>
      </div>
      <div class="grid grid-cols-3"></div>
    </div>
  </div>
</template>
