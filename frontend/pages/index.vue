<script setup lang="ts">
import { calcCareer, formatDate } from "~/util/date";
import { getLanguages } from "~~/api/languages";
import { getCommits } from "~/api/commit";

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

const languages = await getLanguages();
const commits = await getCommits();
</script>

<template>
  <div class="container mx-auto my-5 p-5">
    <div class="md:flex no-wrap md:-mx-2">
      <Profile />
      <!-- Right Side -->
      <div class="w-full md:w-9/12 mx-2 h-64">
        <!-- Profile tab -->
        <!-- About Section -->
        <div class="bg-white p-3 border rounded-lg">
          <div
            class="flex items-center space-x-2 font-semibold text-gray-900 leading-8"
          >
            <span clas="text-green-500">
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
          <div class="text-gray-700">
            <div class="grid md:grid-cols-2 text-sm">
              <div
                class="px-4 py-2 m-2 border rounded-lg hover:border-gray-400"
                @mouseover="viewBarGraph"
              >
                <div class="p-2">Engineer Career</div>
                <div class="p-2 text-2xl font-bold">{{ calcCareer() }}days</div>
                <div class="p-2">
                  {{ viewCareerDate() }}
                </div>
              </div>
              <div
                class="px-4 py-2 m-2 border rounded-lg hover:border-gray-400"
                @mouseover="viewPolarGraph"
                v-if="languages != undefined"
              >
                <div class="p-2">Most Used Language</div>
                <div class="p-2 text-2xl font-bold">
                  {{ languages[0]?.name }}
                </div>
                <div class="p-2">{{ languages[0].ratio }} code</div>
              </div>
            </div>
          </div>
          <div class="h-96" v-if="graph">
            <PolarGraph :languages="languages?.slice(0, 5)" />
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
            <div>
              <div
                class="flex items-center space-x-2 font-semibold text-gray-900 leading-8 mb-3"
              >
                <span clas="text-green-500">
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
                      d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
                    />
                  </svg>
                </span>
                <span class="tracking-wide">Qualifications</span>
              </div>
              <ul class="list-inside space-y-2">
                <li>
                  <a
                    class="text-teal-600"
                    href="https://www.credly.com/earner/earned/badge/8e14359e-5350-4afc-bb2b-f2c9cb1bab87"
                  >
                    AWS Certified Cloud Practitioner
                  </a>
                  <div class="text-gray-500 text-xs">July 2021</div>
                </li>
                <li>
                  <div class="text-teal-600">銀行業務検定法務3級</div>
                  <div class="text-gray-500 text-xs">June 2018</div>
                </li>
                <li>
                  <div class="text-teal-600">日商簿記検定試験3級</div>
                  <div class="text-gray-500 text-xs">February 2018</div>
                </li>
                <li>
                  <div class="text-teal-600">
                    3級ファイナンシャル・プランニング技能士(FP)
                  </div>
                  <div class="text-gray-500 text-xs">October 2017</div>
                </li>
              </ul>
            </div>
            <div>
              <div
                class="flex items-center space-x-2 font-semibold text-gray-900 leading-8 mb-3"
              >
                <span clas="text-green-500">
                  <svg
                    class="h-5"
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path fill="#fff" d="M12 14l9-5-9-5-9 5 9 5z" />
                    <path
                      fill="#fff"
                      d="M12 14l6.16-3.422a12.083 12.083 0 01.665 6.479A11.952 11.952 0 0012 20.055a11.952 11.952 0 00-6.824-2.998 12.078 12.078 0 01.665-6.479L12 14z"
                    />
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M12 14l9-5-9-5-9 5 9 5zm0 0l6.16-3.422a12.083 12.083 0 01.665 6.479A11.952 11.952 0 0012 20.055a11.952 11.952 0 00-6.824-2.998 12.078 12.078 0 01.665-6.479L12 14zm-4 6v-7.5l4-2.222"
                    />
                  </svg>
                </span>
                <span class="tracking-wide">Education and Work Experience</span>
              </div>
              <ul class="list-inside space-y-2">
                <li>
                  <div class="text-teal-600">株式会社Vitalize</div>
                  <div class="text-gray-500 text-xs">February 2021 - Now</div>
                </li>
                <li>
                  <div class="text-teal-600">株式会社福岡銀行</div>
                  <div class="text-gray-500 text-xs">
                    April 2017 - September 2020
                  </div>
                </li>
              </ul>
            </div>
          </div>
          <!-- End of Experience and education grid -->
        </div>
        <!-- End of profile tab -->
      </div>
    </div>
  </div>
</template>
