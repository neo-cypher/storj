// Copyright (C) 2023 Storj Labs, Inc.
// See LICENSE for copying information.

<template>
    <!-- Folder Icon -->
    <svg :width="size" :height="size" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
        <path v-if="!bold" :fill="color" d="M7.41528 1.5L7.47416 1.50022C7.74772 1.5021 7.90752 1.51534 8.08726 1.54897C8.31459 1.59149 8.5204 1.66343 8.72982 1.77453L8.7863 1.80539C8.94923 1.89702 9.10215 2.00659 9.3826 2.23375L11.5641 4.01833L14.7994 4.01839C15.9959 4.023 16.5892 4.13696 17.1909 4.45057L17.2293 4.47085C17.7963 4.7741 18.2443 5.22205 18.5492 5.79226L18.5785 5.84813C18.8702 6.41362 18.9854 6.98502 18.9988 8.07466L19 8.22651V13.3036C18.9951 14.7512 18.8568 15.444 18.4724 16.1629C18.1201 16.8216 17.6002 17.3415 16.9386 17.6953L16.875 17.7287L16.8151 17.7591C16.1749 18.0782 15.5026 18.2062 14.2384 18.2203L14.0773 18.2215H5.99905C4.55062 18.2215 3.8431 18.095 3.13841 17.7354L3.0586 17.6937C2.39991 17.3415 1.88002 16.8216 1.5262 16.16L1.49445 16.0996C1.1502 15.4342 1.01556 14.7633 1.00116 13.4595L1 13.2999V5.25889C1 4.42193 1.1328 3.80687 1.42641 3.25184C1.71281 2.71041 2.13655 2.28105 2.67415 1.98753C3.22526 1.68663 3.83851 1.54572 4.67544 1.53468L7.28265 1.50052L7.41528 1.5ZM9.14621 6.62625C8.90655 6.80284 8.61945 6.90263 8.32276 6.91317L8.27031 6.9141L1.98424 6.91396L1.98438 13.2982C1.98871 14.4884 2.08071 15.0462 2.32046 15.5502L2.33564 15.5816L2.36716 15.6442L2.39578 15.6986C2.65631 16.1858 3.03569 16.5652 3.52283 16.8257C4.09289 17.1306 4.64457 17.2371 5.99905 17.2371H14.0755C15.3047 17.2327 15.8648 17.1328 16.3727 16.8798L16.4229 16.8543L16.4773 16.8257C16.9644 16.5652 17.3438 16.1858 17.6043 15.6986C17.9092 15.1286 18.0157 14.5769 18.0157 13.2224V8.22839C18.0118 7.22984 17.9345 6.77393 17.7342 6.36017L17.7054 6.30266L17.6795 6.25331C17.468 5.85783 17.1605 5.5504 16.7651 5.3389C16.2992 5.08974 15.8481 5.00262 14.7222 5.00262L11.3497 5.00257L9.14621 6.62625ZM7.41761 2.48438L7.29103 2.48486L4.68838 2.51898C4.00304 2.52802 3.54608 2.63302 3.14588 2.85152C2.77781 3.05248 2.49263 3.34144 2.29655 3.71213C2.08334 4.11518 1.98438 4.57348 1.98438 5.25889L1.98424 5.92958L8.27031 5.92972C8.3637 5.92972 8.45492 5.90315 8.53344 5.85347L8.56227 5.83377L10.4911 4.41236L8.839 3.06056C8.59588 2.86164 8.46697 2.76383 8.36672 2.70071L8.35189 2.69149L8.30933 2.66647L8.26352 2.64144C8.14711 2.57971 8.03578 2.54079 7.90626 2.51657C7.80285 2.49722 7.70913 2.48815 7.53803 2.48535L7.41761 2.48438Z" />
        <path v-else :fill="color" d="M7.09273 1.83862C7.2626 1.8704 7.41357 1.92317 7.56625 2.00414L7.61471 2.03057C7.74847 2.1058 7.88213 2.20149 8.19133 2.45447L10.3217 4.19751H13.4957C14.749 4.19751 15.2035 4.328 15.6617 4.57304C16.1198 4.81808 16.4794 5.17766 16.7245 5.63584L16.7507 5.68585C16.9788 6.12813 17.1 6.59407 17.1 7.80178V12.4918C17.1 13.9838 16.9446 14.5249 16.6529 15.0703C16.3612 15.6158 15.9331 16.0439 15.3877 16.3356L15.3331 16.3643L15.2793 16.3916C14.7756 16.6427 14.2223 16.7778 12.8818 16.7826H5.19079C3.69879 16.7826 3.15775 16.6273 2.6123 16.3356C2.06684 16.0439 1.63877 15.6158 1.34705 15.0703L1.31834 15.0158C1.04987 14.4969 0.905038 13.9532 0.899994 12.5656V4.91008C0.899994 4.15367 1.02189 3.65256 1.25157 3.21837C1.48126 2.78417 1.81872 2.44223 2.24985 2.20684C2.68097 1.97145 3.18043 1.84295 3.93678 1.83297L6.41469 1.80048L6.54096 1.79999C6.82689 1.80039 6.95392 1.81266 7.09273 1.83862ZM7.97397 7.07938C7.67373 7.30642 7.31046 7.43381 6.93504 7.44445L6.88378 7.44518H3.41896C3.04556 7.44518 2.6912 7.36369 2.37269 7.21752V12.5604L2.37419 12.7415L2.37599 12.8618C2.39189 13.7187 2.46698 14.0416 2.64573 14.3758C2.80018 14.6646 3.01803 14.8824 3.30683 15.0369L3.34877 15.0588L3.39368 15.081C3.71695 15.2357 4.07688 15.2995 4.95026 15.3087L5.19079 15.3099L12.9325 15.3096C13.9745 15.3045 14.3298 15.2312 14.6932 15.0369C14.982 14.8824 15.1998 14.6646 15.3543 14.3758L15.3762 14.3339L15.3983 14.2889C15.5531 13.9657 15.6168 13.6057 15.6261 12.7324L15.6273 12.4918V7.80178C15.6273 6.88146 15.5746 6.60856 15.4258 6.33038C15.318 6.12885 15.1687 5.97949 14.9671 5.87171L14.9323 5.85361C14.6709 5.72157 14.3902 5.67238 13.5524 5.6703L9.83731 5.67017L7.97397 7.07938ZM6.47489 3.27269L6.34265 3.27408L3.95621 3.30557L3.84915 3.30796C3.41773 3.32185 3.16585 3.38465 2.95559 3.49945C2.77811 3.59635 2.64794 3.72825 2.55338 3.907C2.4389 4.12343 2.3793 4.38262 2.37324 4.83186C2.37668 4.85579 2.37883 4.87987 2.37979 4.90425L2.38037 4.93387C2.38037 5.49643 2.82765 5.95453 3.38599 5.97194L3.41896 5.97245H6.88378C6.94754 5.97245 7.00978 5.95424 7.06329 5.92024L7.08566 5.90471L8.64293 4.72696L7.17963 3.52982C7.00245 3.38612 6.93952 3.34067 6.88962 3.31253L6.87625 3.30522C6.8551 3.29401 6.84547 3.29064 6.82193 3.28624L6.79088 3.28128L6.76967 3.27886C6.71675 3.27363 6.63713 3.27155 6.47489 3.27269Z" />
    </svg>
</template>

<script setup lang="ts">
const props = withDefaults(defineProps<{
    size?: number | string;
    color?: string;
    bold?: boolean;
}>(), {
    size: 16,
    color: 'currentColor',
    bold: false,
});
</script>
