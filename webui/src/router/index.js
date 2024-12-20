import {createRouter, createWebHashHistory} from 'vue-router'
import Test from '../views/ConversationsParentView.vue'
import LoginView from '../views/LoginView.vue'
import ConversationsParentView from '../views/ConversationsParentView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: Test},
		{path: '/login', component: LoginView},
		{path: '/conversations', component: ConversationsParentView},
		// {path: '/some/:id/link', component: HomeView},
	]
})

export default router
