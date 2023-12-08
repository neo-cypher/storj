// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

import { shallowMount } from '@vue/test-utils';

import { ProjectMemberHeaderState } from '@/types/projectMembers';

import HeaderArea from '@/components/team/HeaderArea.vue';

describe('Team HeaderArea', () => {
    it('renders correctly', () => {
        const wrapper = shallowMount(HeaderArea);

        const addNewTemMemberPopup = wrapper.findAll('adduserpopup-stub');

        expect(wrapper).toMatchSnapshot();
        expect(addNewTemMemberPopup.length).toBe(0);
        expect(wrapper.findAll('.header-default-state').length).toBe(1);
        expect(wrapper.findAll('.blur-content').length).toBe(0);
        expect(wrapper.findAll('.blur-search').length).toBe(0);
        expect(wrapper.vm.isDeleteClicked).toBe(false);
    });

    it('renders correctly with opened Add team member popup', () => {
        const wrapper = shallowMount(HeaderArea);

        expect(wrapper).toMatchSnapshot();
        expect(wrapper.findAll('.header-default-state').length).toBe(1);
        expect(wrapper.vm.isDeleteClicked).toBe(false);
        expect(wrapper.findAll('.blur-content').length).toBe(0);
        expect(wrapper.findAll('.blur-search').length).toBe(0);
    });

    it('renders correctly with selected users', () => {
        const selectedUsersCount = 2;

        const wrapper = shallowMount(HeaderArea, {
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
        const selectedUsersCount = 2;

        const wrapper = shallowMount(HeaderArea, {
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
        const selectedUsersCount = 1;

        const wrapper = shallowMount(HeaderArea, {
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
