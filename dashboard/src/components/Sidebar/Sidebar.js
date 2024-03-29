import React from "react";
import classNames from "classnames";
import { logout } from "../../actions/user";
import { useDispatch } from "react-redux";
import { NavLink, Route } from "react-router-dom";
import routes from "../../routes/sidebar";
import { RiLogoutBoxFill } from "react-icons/ri";
import {
  ChevronDoubleLeftIcon,
  ChevronDoubleRightIcon,
} from "@heroicons/react/24/outline";

const Sidebar = ({ collapsed, shown, setCollapsed }) => {
  const dispatch = useDispatch();
  const Icon = collapsed ? ChevronDoubleRightIcon : ChevronDoubleLeftIcon;

  const handleLogout = () => {
    dispatch(logout());
  };

  return (
    <div
      className={classNames({
        "bg-sky-800 text-zinc-50 fixed md:static md:translate-x-0 z-20": true,
        "transition-all duration-300 ease-in-out": true,
        "w-[300px]": !collapsed,
        "w-16": collapsed,
        "-translate-x-full": !shown,
      })}
      style={{ height: "100%" }}
    >
      <div
        className={classNames({
          "flex flex-col justify-between sticky inset-0 w-full": true,
        })}
      >
        {/* logo and collapse button */}
        <div
          className={classNames({
            "flex items-center transition-none": true,
            "p-4 justify-between": !collapsed,
            "py-4 justify-center": collapsed,
          })}
        >
          {!collapsed && (
            <div className="flex items-center space-x-4">
              <img
                src={require("../../assets/img/logo_white.png")}
                height={36}
                width={36}
                alt="Logo"
                className="rounded-full"
              />
              <div className="text-xl font-semibold">Rubyx</div>
            </div>
          )}
          <button
            className="grid place-content-center hover:bg-sky-900 w-10 h-10 rounded-full opacity-0 md:opacity-100"
            onClick={() => setCollapsed(!collapsed)}
          >
            <Icon className="w-5 h-5" />
          </button>
        </div>
        <nav className="flex-grow">
          <ul
            className={classNames({
              "my-2 flex flex-col gap-2 items-stretch": true,
            })}
          >
            {routes.map((item, index) => {
              return (
                <li
                  key={index}
                  className={classNames({
                    "text-indigo-100 hover:bg-sky-900 flex": true, //colors
                    "transition-colors duration-300": true, //animation
                    "rounded-md p-2 mx-3 gap-4 ": !collapsed,
                    "rounded-full p-2 mx-3 w-10 h-10": collapsed,
                  })}
                >
                  <NavLink
                    exact
                    to={item.href}
                    className="flex gap-2"
                    activeClassName="text-white dark:text-gray-100"
                  >
                    <Route path={item.href}>
                      <span
                        className="absolute inset-y-0 left-0 w-1  rounded-tr-lg rounded-br-lg"
                        aria-hidden="true"
                      ></span>
                    </Route>
                    {item.icon} <span>{!collapsed && item.label}</span>
                  </NavLink>
                </li>
              );
            })}
            <li
              className={classNames({
                "text-indigo-100 hover:bg-sky-900 flex cursor-pointer": true, //colors
                "transition-colors duration-300": true, //animation
                "rounded-md p-2 mx-3 gap-4 ": !collapsed,
                "rounded-full p-2 mx-3 w-10 h-10": collapsed,
              })}
              onClick={handleLogout}
            >
              <div className="flex gap-2">
                <span
                  className="absolute inset-y-0 left-0 w-1  rounded-tr-lg rounded-br-lg"
                  aria-hidden="true"
                ></span>
                <RiLogoutBoxFill className="w-6 h-6" />{" "}
                <span>{!collapsed && "Log Out"}</span>
              </div>
            </li>
          </ul>
        </nav>
      </div>
    </div>
  );
};
export default Sidebar;
