import { Fragment as ReactFragment } from "react";

import type { ComponentRegistry, RegistryEntry } from "./types";
import {
  ButtonBuiltin,
  CardBuiltin,
  CheckboxBuiltin,
  ColumnBuiltin,
  ContainerBuiltin,
  HeadingBuiltin,
  InputBuiltin,
  RowBuiltin,
  TextBuiltin,
} from "./widgets";

function entry(component: RegistryEntry["component"]): RegistryEntry {
  return { component };
}

/** Framework owned widgets — no generated kit registry required. */
export const builtinRegistry: ComponentRegistry = {
  Fragment: entry(ReactFragment as RegistryEntry["component"]),
  Text: entry(TextBuiltin),
  Heading: entry(HeadingBuiltin),
  Row: entry(RowBuiltin),
  Column: entry(ColumnBuiltin),
  Container: entry(ContainerBuiltin),
  Card: entry(CardBuiltin),
  Button: entry(ButtonBuiltin),
  Input: entry(InputBuiltin),
  Checkbox: entry(CheckboxBuiltin),
};
