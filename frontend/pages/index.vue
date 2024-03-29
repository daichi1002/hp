<script setup lang="ts">
import { calcCareer, formatDate } from "~/util/date";

const viewCareerDate = () => {
  const startDate = formatDate(new Date("2021-02-08"));
  const today = formatDate(new Date());

  return `${startDate}~${today}`;
};

const graph = ref(true);
const viewPolarGraph = () => {
  graph.value = true;
};
const viewBarGraph = () => {
  graph.value = false;
};

const { languages, commits, getLanguages, getCommits } = useContributionStore();
onMounted(() => getLanguages);
onMounted(() => getCommits);
</script>

<template>
  <div class="container mx-auto my-5 p-5">
    <div class="md:flex no-wrap md:-mx-2">
      <Profile />
      <!-- Right Side -->
      <div class="w-full md:w-9/12 mx-2 h-max">
        <!-- Profile tab -->
        <!-- About Section -->
        <div class="bg-white p-3 border rounded-lg">
          <div class="flex items-center">
            <div
              class="flex items-center space-x-2 font-semibold text-gray-900 leading-8"
            >
              <span>
                <svg
                  class="h-5"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"
                  />
                </svg>
              </span>
              <span class="tracking-wide">Engineer Profile</span>
            </div>
          </div>
          <div class="text-gray-700">
            <div class="grid md:grid-cols-2 text-sm">
              <div
                class="px-4 py-2 m-4 border rounded-lg updown1"
                @mouseover="viewBarGraph"
              >
                <div class="p-2">Engineer Career</div>
                <div class="p-2 text-2xl font-bold">{{ calcCareer() }}days</div>
                <div class="p-2">
                  {{ viewCareerDate() }}
                </div>
              </div>
              <div
                class="px-4 py-2 m-4 border rounded-lg updown2"
                @mouseover="viewPolarGraph"
              >
                <div class="p-2">Most Used Language</div>
                <div class="p-2 text-2xl font-bold">
                  {{ languages[0].name }}
                </div>
                <div class="p-2">
                  {{ languages[0].ratio.toLocaleString() }} code
                </div>
              </div>
            </div>
          </div>
          <div class="h-96" v-if="graph">
            <PolarGraph :languages="languages.slice(0, 5)" />
          </div>
          <div class="h-96" v-else>
            <BarGraph :commits="commits?.slice(-7)" />
          </div>
        </div>
        <!-- End of about section -->
        <div class="my-4"></div>
        <!-- Experience and education -->
        <div class="bg-white p-3 border rounded-lg">
          <div class="grid grid-cols-2">
            <Qualifications />
            <PersonalHistory />
          </div>
          <!-- End of Experience and education grid -->
        </div>
        <!-- End of profile tab -->
      </div>
    </div>
  </div>
</template>

<style>
.updown1 {
  animation-name: updown1; /* アニメーション名の指定 */
  animation-delay: 0s; /* アニメーションの開始時間指定 */
  animation-duration: 6s; /* アニメーション動作時間の指定 */
  animation-timing-function: ease-in-out;
  /* アニメーションの動き（徐々に早く徐々に遅く）*/
  animation-iteration-count: infinite; /* アニメーションをループさせる */
}

.updown2 {
  animation-name: updown1; /* アニメーション名の指定 */
  animation-delay: 3s; /* アニメーションの開始時間指定 */
  animation-duration: 6s; /* アニメーション動作時間の指定 */
  animation-timing-function: ease-in-out;
  /* アニメーションの動き（徐々に早く徐々に遅く）*/
  animation-iteration-count: infinite; /* アニメーションをループさせる */
}

@keyframes updown1 {
  0% {
    transform: scale(1);
  }

  40% {
    transform: scale(1);
  }

  50% {
    transform: scale(1.1);
  }

  60% {
    transform: scale(1);
  }

  100% {
    transform: scale(1);
  }
}
</style>
