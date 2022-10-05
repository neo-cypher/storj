// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

import Vuex from 'vuex';
import { createLocalVue, shallowMount } from '@vue/test-utils';

import { ProjectMembersApiMock } from '../mock/api/projectMembers';

import { appStateModule } from '@/store/modules/appState';
import { makeNotificationsModule } from '@/store/modules/notifications';
import { makeProjectMembersModule } from '@/store/modules/projectMembers';
import { ProjectMember, ProjectMemberHeaderState, ProjectMembersPage } from '@/types/projectMembers';
import { APP_STATE_MUTATIONS } from '@/store/mutationConstants';

import HeaderArea from '@/components/team/HeaderArea.vue';

const localVue = createLocalVue();
const date = new Date(0);
const projectMembers: ProjectMember[] = [
    new ProjectMember('f1', 's1', '1@example.com', date, '1'),
    new ProjectMember('f2', 's2', '2@example.com', date, '2'),
    new ProjectMember('f3', 's3', '3@example.com', date, '3'),
];

const api = new ProjectMembersApiMock();
api.setMockPage(new ProjectMembersPage(projectMembers));

const notificationsModule = makeNotificationsModule();
const projectMembersModule = makeProjectMembersModule(api);

localVue.use(Vuex);

const store = new Vuex.Store({ modules: { appStateModule, projectMembersModule, notificationsModule } });

describe('Team HeaderArea', () => {
    it('renders correctly', () => {
        const wrapper = shallowMount<HeaderArea>(HeaderArea, {
            store,
            localVue,
        });

        const addNewTemMemberPopup = wrapper.findAll('adduserpopup-stub');

        expect(wrapper).toMatchSnapshot();
        expect(addNewTemMemberPopup.length).toBe(0);
        expect(wrapper.findAll('.header-default-state').length).toBe(1);
        expect(wrapper.findAll('.blur-content').length).toBe(0);
        expect(wrapper.findAll('.blur-search').length).toBe(0);
        expect(wrapper.vm.isDeleteClicked).toBe(false);
    });

    it('renders correctly with opened Add team member popup', () => {
        store.commit(APP_STATE_MUTATIONS.TOGGLE_ADD_TEAM_MEMBERS_MODAL);

        const wrapper = shallowMount<HeaderArea>(HeaderArea, {
            store,
            localVue,
        });

        expect(wrapper).toMatchSnapshot();
        expect(wrapper.findAll('.header-default-state').length).toBe(1);
        expect(wrapper.vm.isDeleteClicked).toBe(false);
        expect(wrapper.findAll('.blur-content').length).toBe(0);
        expect(wrapper.findAll('.blur-search').length).toBe(0);

        store.commit(APP_STATE_MUTATIONS.TOGGLE_ADD_TEAM_MEMBERS_MODAL);
    });

    it('renders correctly with selected users', () => {
        store.commit(APP_STATE_MUTATIONS.TOGGLE_ADD_TEAM_MEMBERS_MODAL);

        const selectedUsersCount = 2;

        const wrapper = shallowMount<HeaderArea>(HeaderArea, {
            store,
            localVue,
            propsData: {
                selectedProjectMembersCount: selectedUsersCount,
                headerState: ProjectMemberHeaderState.ON_SELECT,
            },
        });

        expect(wrapper.findAll('.header-selected-members').length).toBe(1);

        expect(wrapper).toMatchSnapshot();
        expect(wrapper.vm.selectedProjectMembersCount).toBe(selectedUsersCount);
        expect(wrapper.vm.isDeleteClicked).toBe(false);
        expect(wrapper.findAll('.blur-content').length).toBe(0);
        expect(wrapper.findAll('.blur-search').length).toBe(0);
    });

    it('renders correctly with 2 selected users and delete clicked once', async () => {
        store.commit(APP_STATE_MUTATIONS.TOGGLE_ADD_TEAM_MEMBERS_MODAL);

        const selectedUsersCount = 2;

        const wrapper = shallowMount<HeaderArea>(HeaderArea, {
            store,
            localVue,
            propsData: {
                selectedProjectMembersCount: selectedUsersCount,
                headerState: ProjectMemberHeaderState.ON_SELECT,
            },
        });

        await wrapper.vm.onFirstDeleteClick();

        expect(wrapper).toMatchSnapshot();
        expect(wrapper.vm.selectedProjectMembersCount).toBe(selectedUsersCount);
        expect(wrapper.vm.userCountTitle).toBe('users');
        expect(wrapper.vm.isDeleteClicked).toBe(true);
        expect(wrapper.findAll('.blur-content').length).toBe(1);
        expect(wrapper.findAll('.blur-search').length).toBe(1);

        const expectedSectionRendered = wrapper.find('.header-after-delete-click');
        expect(expectedSectionRendered.text()).toBe(`Are you sure you want to delete ${selectedUsersCount} users?`);
    });

    it('renders correctly with 1 selected user and delete clicked once', async () => {
        store.commit(APP_STATE_MUTATIONS.TOGGLE_ADD_TEAM_MEMBERS_MODAL);

        const selectedUsersCount = 1;

        const wrapper = shallowMount<HeaderArea>(HeaderArea, {
            store,
            localVue,
            propsData: {
                selectedProjectMembersCount: selectedUsersCount,
                headerState: ProjectMemberHeaderState.ON_SELECT,
            },
        });

        await wrapper.vm.onFirstDeleteClick();

        expect(wrapper).toMatchSnapshot();
        expect(wrapper.vm.selectedProjectMembersCount).toBe(selectedUsersCount);
        expect(wrapper.vm.userCountTitle).toBe('user');
        expect(wrapper.vm.isDeleteClicked).toBe(true);
        expect(wrapper.findAll('.blur-content').length).toBe(1);
        expect(wrapper.findAll('.blur-search').length).toBe(1);

        const expectedSectionRendered = wrapper.find('.header-after-delete-click');
        expect(expectedSectionRendered.text()).toBe(`Are you sure you want to delete ${selectedUsersCount} user?`);
    });
});
