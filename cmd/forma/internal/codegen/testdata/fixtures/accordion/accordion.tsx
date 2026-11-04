import * as React from "react";

export type AccordionProps = {
  className?: string;
  children?: React.ReactNode;
};

export function Accordion({ className, children }: AccordionProps) {
  return <div className={className}>{children}</div>;
}

export function AccordionItem({
  className,
  children,
}: {
  className?: string;
  children?: React.ReactNode;
}) {
  return <div className={className}>{children}</div>;
}

export function AccordionTrigger({
  className,
  children,
}: {
  className?: string;
  children?: React.ReactNode;
}) {
  return (
    <button type="button" className={className}>
      {children}
    </button>
  );
}

export function AccordionContent({
  className,
  children,
}: {
  className?: string;
  children?: React.ReactNode;
}) {
  return <div className={className}>{children}</div>;
}
