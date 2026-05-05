// /cluster/dashboard
export const clusterRoutes = {
  path: '/cluster',
  component: () => import('../view/layout/Layout.vue'),
  children: [
    {
      path: 'dashboard',
      component: () => import('../view/dashboard/Dashboard.vue'),
    },
    {
      path: 'list',
      component: () => import('../view/cluster/cluster.vue'),
    },
    {
      path: 'events',
      component: () => import('../view/event/Events.vue'),
    },
  ],
}
