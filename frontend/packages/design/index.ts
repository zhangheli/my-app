import type { App, Plugin } from "vue";
import "@jinyaoma/my-app-components/theme-default/style.scss";
import { MyIcon, MyContainer } from "@jinyaoma/my-app-components";

const components = [MyIcon, MyContainer] as Plugin[];

const install = (app: App) => {
  components.map((item) => {
    app.use(item);
  });
};

export default {
  install,
} as Plugin;
