// Copyright (C) 2024 Storj Labs, Inc.
// See LICENSE for copying information.

<template>
    <v-dialog
        v-model="model"
        activator="parent"
        width="auto"
        min-width="400px"
        max-width="450px"
        transition="fade-transition"
    >
        <v-card rounded="xlg">
            <v-sheet>
                <v-card-item class="py-4 pl-6">
                    <template #prepend>
                        <img src="@/assets/icon-color-globe.svg" alt="Earth" width="40" class="mt-1">
                    </template>
                    <v-card-title class="font-weight-bold">
                        Storj Sustainability
                    </v-card-title>
                    <template #append>
                        <v-btn
                            icon="$close"
                            variant="text"
                            size="small"
                            color="default"
                            @click="model = false"
                        />
                    </template>
                </v-card-item>
            </v-sheet>

            <v-card-text class="mt-n4 px-6">
                <v-card class="pa-4 mb-4">
                    <p class="text-body-2 font-weight-bold mb-2">Carbon Emissions</p>
                    <v-chip variant="tonal" color="info" class="font-weight-bold">
                        {{ emission.storjImpact.toLocaleString(undefined, { maximumFractionDigits: 0 }) }} kg CO₂e
                    </v-chip>
                    <p class="text-body-2 mt-2">Estimated for this Storj project. <a href="https://www.storj.io/documents/storj-sustainability-whitepaper.pdf" target="_blank" rel="noopener noreferrer" class="link font-weight-bold">Learn more</a></p>
                </v-card>
                <v-card class="pa-4 mb-4">
                    <p class="text-body-2 font-weight-bold mb-2">Carbon Comparison</p>
                    <v-chip variant="tonal" color="warning" class="font-weight-bold">
                        {{ emission.hyperscalerImpact.toLocaleString(undefined, { maximumFractionDigits: 0 }) }} kg CO₂e
                    </v-chip>
                    <p class="text-body-2 mt-2">By using traditional cloud storage. <a href="https://www.storj.io/documents/storj-sustainability-whitepaper.pdf" target="_blank" rel="noopener noreferrer" class="link font-weight-bold">Learn more</a></p>
                </v-card>
                <v-card class="pa-4 mb-4">
                    <p class="text-body-2 font-weight-bold mb-2">Total Carbon Avoided</p>
                    <v-chip variant="tonal" color="green" class="font-weight-bold">
                        {{ co2Savings }} kg CO₂e
                    </v-chip>
                    <p class="text-body-2 mt-2">Estimated by using Storj. <a href="https://www.storj.io/documents/storj-sustainability-whitepaper.pdf" target="_blank" rel="noopener noreferrer" class="link font-weight-bold">Learn more</a></p>
                </v-card>
                <v-card class="pa-4 mb-2">
                    <p class="text-body-2 font-weight-bold mb-2">Carbon Avoided Equals To</p>
                    <v-chip variant="tonal" color="green" class="font-weight-bold">
                        {{ emission.savedTrees.toLocaleString() }} trees grown for 10 years
                    </v-chip>
                    <p class="text-body-2 mt-2">Estimated equivalencies. <a href="https://www.epa.gov/energy/greenhouse-gases-equivalencies-calculator-calculations-and-references#seedlings" target="_blank" rel="noopener noreferrer" class="link font-weight-bold">Learn more</a></p>
                </v-card>
            </v-card-text>

            <v-card-actions class="pa-6 pt-2">
                <v-row>
                    <v-col>
                        <v-btn color="primary" variant="flat" block link href="https://www.storj.io/documents/storj-sustainability-whitepaper.pdf" target="_blank" rel="noopener noreferrer">
                            Sustainability Whitepaper <v-icon :icon="mdiOpenInNew" class="ml-2" />
                        </v-btn>
                    </v-col>
                </v-row>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';
import {
    VDialog,
    VCard,
    VSheet,
    VCardItem,
    VCardActions,
    VCardTitle,
    VCardText,
    VBtn,
    VChip,
    VRow,
    VCol,
    VIcon,
} from 'vuetify/components';
import { mdiOpenInNew } from '@mdi/js';

import { Emission } from '@/types/projects';
import { useProjectsStore } from '@/store/modules/projectsStore';

const projectsStore = useProjectsStore();

const model = ref<boolean>(false);

/**
 * Returns project's emission impact.
 */
const emission = computed<Emission>(()  => {
    return projectsStore.state.emission;
});

/**
 * Returns calculated and formatted CO2 savings info.
 */
const co2Savings = computed<string>(() => {
    let saved = emission.value.hyperscalerImpact - emission.value.storjImpact;
    if (saved < 0) saved = 0;

    return saved.toLocaleString(undefined, { maximumFractionDigits: 0 });
});
</script>
