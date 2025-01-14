import {createRouter, createWebHashHistory} from 'vue-router'
import LoginView from '../views/LoginView.vue'
import NewConversationsParentView from '../views/NewConversationsParentView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/login', component: LoginView},
		{path: '/conversations', component: NewConversationsParentView},
		{path: '/conversation/:id', component: NewConversationsParentView},
		// {path: '/some/:id/link', component: HomeView},
	]
})

export default router
